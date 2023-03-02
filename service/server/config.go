package server

import "time"

type Config struct {
	Addr       string
	CertFile   string
	KeyFile    string
	SessionTTL time.Duration
}
