package client

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/karpovicham/secured-tcp-server/internal/logger"
	"github.com/karpovicham/secured-tcp-server/internal/messenger"
	"github.com/karpovicham/secured-tcp-server/service/client/resolver"
	"io"
	"net"
	"sync"
)

type Client struct {
	Log logger.Logger
	Cfg Config
}

func NewClient(cfg Config, log logger.Logger) *Client {
	return &Client{
		Cfg: cfg,
		Log: log,
	}
}

// Run connects to the server
// and start getting Quote every 3 seconds in endless loop
func (c *Client) Run(ctx context.Context) error {
	// Certificated generated with openssl to establish secure (TLS) connection
	cert, err := tls.LoadX509KeyPair(c.Cfg.CertFile, c.Cfg.KeyFile)
	if err != nil {
		return fmt.Errorf("load key pair: %w", err)
	}

	// Create a socket and start listening
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", c.Cfg.Addr, &config)
	if err != nil {
		return fmt.Errorf("deal: %w", err)
	}
	defer conn.Close()

	c.Log.Info("Client", conn.LocalAddr(), "connected to:", conn.RemoteAddr())

	// Track context status
	go watchContext(ctx, conn)

	apiResolver := resolver.NewClientAPIResolver(messenger.NewMessenger(conn))

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Endless loop to get a new Quote every 3 secs
		//for {
		if err = c.DemoProcess(ctx, apiResolver); err != nil {
			switch true {
			case contextCanceled(ctx):
			case errors.Is(err, io.EOF):
				c.Log.Error("Server closed connection")
			default:
				c.Log.Error("DemoProcess quote: ", err)
			}
			return
		}

		//time.Sleep(3 * time.Second)
		//}
	}()

	wg.Wait()
	return nil
}

// Context could be closed by manual interrupt.
// Close listener connection to stop serving connections and as the result - service runtime.
func watchContext(ctx context.Context, conn net.Conn) {
	<-ctx.Done()
	conn.Close()
}

// contextCanceled returns true if the context is canceled
func contextCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
