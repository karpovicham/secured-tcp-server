package server

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/karpovicham/secured-tcp-server/internal/logger"
	"github.com/karpovicham/secured-tcp-server/internal/messenger"
	"github.com/karpovicham/secured-tcp-server/internal/proto"
)

type Server struct {
	Log logger.Logger
	Cfg Config
	// Messenger for sending and receiving messages over the TCP connection on the protocol level
	MsgrFn func(conn net.Conn) messenger.Messenger
	// To work with DB
	Repo Repo
	// To mock time for tests
	NowFn func() time.Time
}

func NewTCPServer(cfg Config, logger logger.Logger, repo Repo, msgrFn messenger.MsgrFn) *Server {
	return &Server{
		Cfg:    cfg,
		Log:    logger,
		MsgrFn: msgrFn,
		Repo:   repo,
		NowFn: func() time.Time {
			return time.Now().UTC()
		},
	}
}

// Run - start the server and serve new client connection in endless loop
func (s *Server) Run(ctx context.Context) error {
	// Certificated generated with openssl to establish secure (TLS) connection
	// Better to use services like Consul to store secured data
	cert, err := tls.LoadX509KeyPair(s.Cfg.CertFile, s.Cfg.KeyFile)
	if err != nil {
		return fmt.Errorf("load key pair: %w", err)
	}

	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAnyClientCert,
		Rand:         rand.Reader,
	}

	// Init server listener
	ln, err := tls.Listen("tcp", s.Cfg.Addr, &config)
	if err != nil {
		return err
	}
	defer ln.Close()

	s.Log.Info("Listening:", ln.Addr())

	// Track context status
	go watchContext(ctx, ln)

	// Us WG is gracefully shut down the serving process
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			if contextCanceled(ctx) {
				return
			}

			conn, err := ln.Accept()
			if err != nil {
				switch true {
				// Error could be received for closed connection by manual interrupt
				case contextCanceled(ctx):
					return
				default:
					s.Log.Error("Accepting conn:", err)
					continue
				}
			}

			// Effectively this should be implemented with a workers pool
			go s.handleConnection(ctx, conn)
		}
	}()

	// Waiting for server to stop serving connections
	wg.Wait()
	return nil
}

// Context could be closed by manual interrupt.
// Close listener connection to stop serving connections and as the result - service runtime.
func watchContext(ctx context.Context, ln net.Listener) {
	<-ctx.Done()
	ln.Close()
}

// Serve client connection
func (s *Server) handleConnection(ctx context.Context, conn net.Conn) {
	clientAddr := conn.RemoteAddr()
	s.Log.Info("New client:", clientAddr.String())
	defer s.Log.Info("Close client:", clientAddr.String())
	defer conn.Close()

	handler := NewRequestsHandler(s, conn)
	for {
		receivedMsg, err := handler.Msgr.Receive()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			s.Log.Error("Receive:", err)
			return
		}

		switch receivedMsg.Type {
		case proto.TypeLogin:
			if err = handler.HandleLoginRequest(ctx, receivedMsg); err != nil {
				s.Log.Error("HandleLoginRequest:", err)
				return
			}
		case proto.TypeModifyAccountSettings:
			if err = handler.HandleModifyAccountSettingsRequest(ctx, receivedMsg); err != nil {
				s.Log.Error("HandleModifyAccountSettingsRequest:", err)
				return
			}
		case proto.TypeDeactivateUser:
			if err = handler.HandleDeactivateUserRequest(ctx, receivedMsg); err != nil {
				s.Log.Error("HandleDeactivateUserRequest:", err)
				return
			}
		case proto.TypeAddFavoritePage:
			if err = handler.HandleAddFavoritePageRequest(ctx, receivedMsg); err != nil {
				s.Log.Error("HandleAddFavoritePageRequest:", err)
				return
			}
		case proto.TypeLogout:
			if err = handler.HandleLogoutRequest(ctx, receivedMsg); err != nil {
				s.Log.Error("HandleLogoutRequest:", err)
				return
			}
		default:
			s.Log.Error("Unsupported protocols:", receivedMsg.Type, err)
			return
		}
	}
}

// contextCanceled returns true if the context is contextCanceled
func contextCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
