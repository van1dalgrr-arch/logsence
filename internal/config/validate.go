package config

import (
	"fmt"
)

func (c *Config) Validate() error {
	if c.Http.Port < 1 || c.Http.Port > 65535 {
		return fmt.Errorf("invalid port: %d", c.Http.Port)
	}
	if c.Database.Port < 1 || c.Database.Port > 65535 {
		return fmt.Errorf("invalid database port: %d", c.Database.Port)
	}
	if c.Log.Level != "debug" && c.Log.Level != "info" && c.Log.Level != "warn" && c.Log.Level != "error" {
		return fmt.Errorf("invalid log level: %s", c.Log.Level)
	}
	if c.Log.Format != "json" && c.Log.Format != "text" {
		return fmt.Errorf("invalid log format: %s", c.Log.Format)
	}
	return nil
}
