package main

import (
	"context"
	"flag"
	"github.com/jackc/pgx/v4"
	"github.com/karpovicham/secured-tcp-server/internal/logger"
	"github.com/karpovicham/secured-tcp-server/internal/messenger"
	"github.com/karpovicham/secured-tcp-server/internal/signal"
	"github.com/karpovicham/secured-tcp-server/service/server"
	"github.com/karpovicham/secured-tcp-server/service/server/db"
	"os"
	"time"
)

var (
	addr       = flag.String("addr", ":9992", "Connection host and port")
	certFile   = flag.String("cert_file", "etc/server.pem", "Path to certificate file for secured connection")
	keyFile    = flag.String("key_file", "etc/server.key", "Path to key file for secured connection")
	sessionTTL = flag.Duration("session_ttl", 30*time.Second, "TLL of the session ID generated for the client")
	pgURL      = flag.String("pg_url", "postgres://windscribe:windscribe@db:5432/windscribe", "URL config for PostgreSQL connection")
)

func main() {
	// Read cmd arguments
	flag.Parse()
	cfg := server.Config{
		Addr:       *addr,
		CertFile:   *certFile,
		KeyFile:    *keyFile,
		SessionTTL: *sessionTTL,
	}

	log := logger.NewLogger(os.Stdout)

	ctx, cancel := context.WithCancel(context.Background())
	go signal.WatchShutdown(cancel)

	// Better to customize for production code, ie keep alive, max connections etc configs
	conn, err := pgx.Connect(ctx, *pgURL)
	if err != nil {
		log.Fatal("Unable to connect to database: %w", err)
	}
	defer conn.Close(ctx)

	s := server.NewTCPServer(cfg, log, db.NewRepo(conn), messenger.NewMessenger)
	if err := s.Run(ctx); err != nil {
		log.Fatal("Run server:", err)
	}
}
