package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config struct to read file
type Config struct {
	Server  string
	LogMode bool
}

//Read func
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
