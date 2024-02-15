package config

import "time"

const (
	ReadTimeout                   = 5 * time.Second
	WriteTimeout                  = 10 * time.Second
	IdleTimeout                   = 15 * time.Second
	GracefulShutdownTimeout       = 5 * time.Second
	PreflightCacheDurationSeconds = 300
	DefaultUUID                   = "00000000-0000-0000-0000-000000000000"
)

type ContextKey string

const ContextUserIDKey ContextKey = "userID"
