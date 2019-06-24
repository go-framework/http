package server

type Config struct {
	// TCP address to listen on, ":http" if empty
	Addr string `json:"addr" yaml:"addr"`

	CertFile string `json:"cert_file" yaml:"cert_file"`
	KeyFile  string `json:"key_file" yaml:"key_file"`
}

func GetDefaultConfig() *Config {
	return &Config{
		Addr: ":http",
	}
}
