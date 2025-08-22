package agentx

import (
	"log/slog"
	"time"
)

type dialOptions struct {
	logger            *slog.Logger
	timeout           time.Duration
	reconnectInterval time.Duration
}

type DialOption func(o *dialOptions)

func WithLogger(value *slog.Logger) DialOption {
	return func(o *dialOptions) {
		o.logger = value
	}
}

func WithTimeout(value time.Duration) DialOption {
	return func(o *dialOptions) {
		o.timeout = value
	}
}

func WithReconnectInterval(value time.Duration) DialOption {
	return func(o *dialOptions) {
		o.reconnectInterval = value
	}
}
