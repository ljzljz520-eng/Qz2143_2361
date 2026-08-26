package config

import "os"

type Config struct {
	DataPath string
	Addr     string
}

func Load() Config {
	p := os.Getenv("COMMUNITY50_DB")
	if p == "" {
		p = "community50.db"
	}
	a := os.Getenv("COMMUNITY50_ADDR")
	if a == "" {
		a = ":8080"
	}
	return Config{DataPath: p, Addr: a}
}
func (c Config) Secure() bool { return c.DataPath != "" && c.Addr != "" }
