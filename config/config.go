package config

// Config is config!
type Config struct {
	HTTPAddr string `usage:"HTTP server address"`
	Statics  string `usage:"Static files dir"`
}
