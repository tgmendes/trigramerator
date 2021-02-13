package config

import "time"

// Config contains all configurations for the app.
// There will be a bit of "stuttering" when using this package (i.e. config.Config).
// This is an acceptable tradeoff in this context.
type Config struct {
	Server Server
}

// Server contains all server related configurations
type Server struct {
	Host string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	ShutdownTimeout time.Duration
}
