package server

import "time"

// Config consists of informations about the Server.
// This should be a pure struct which doesn't contain instances
// of other structs so that it can be serialized easily.
type Config struct {
	// The address which the server will listen on
	Addr string

	// The path of assets directory. If it is a http url, the server
	// will proxy the asset requests to the url.
	AssetsPath string

	// The secret key that will be used when generating tokens
	SecretKey string

	// The timeout for graceful shutdowns in seconds.
	GracefulTimeout time.Duration
}
