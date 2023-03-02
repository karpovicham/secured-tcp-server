package main

import (
	"context"
	"flag"
	"github.com/karpovicham/secured-tcp-server/internal/logger"
	"github.com/karpovicham/secured-tcp-server/internal/signal"
	"github.com/karpovicham/secured-tcp-server/service/client"
	"os"
)

var (
	addr     = flag.String("addr", ":9992", "Connection host and port")
	certFile = flag.String("cert_file", "etc/client.pem", "Path to certificate file for secured connection")
	keyFile  = flag.String("key_file", "etc/client.key", "Path to key file for secured connection")
)

func main() {
	// Read cmd arguments
	flag.Parse()
	cfg := client.Config{
		Addr:     *addr,
		CertFile: *certFile,
		KeyFile:  *keyFile,
	}

	log := logger.NewLogger(os.Stdout)

	ctx, cancel := context.WithCancel(context.Background())
	go signal.WatchShutdown(cancel)

	c := client.NewClient(cfg, log)
	if err := c.Run(ctx); err != nil {
		log.Fatal("Run client:", err)
	}
}
