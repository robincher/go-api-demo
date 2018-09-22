package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type DBConfig struct {
	Server   string
	Database string
}

// Read and parse the MongoDB configuration file
func (c *DBConfig) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
