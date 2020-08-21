package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//Config
type Config struct {
	Server  string
	LogMode bool
}

//Read function to read a file and get your values
func (c *Config) Read() {

	var fpath = "config.toml"

	if os.Getenv("FILE_PATH_TOML") != "" {
		fpath = os.Getenv("FILE_PATH_TOML")
	}

	if _, err := toml.DecodeFile(fpath, &c); err != nil {
		log.Fatal(err)
	}
}
