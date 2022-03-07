package config

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

//Config - config for application
type Config struct {
	Environment string
	Server      struct {
		Host string
		Port string
	}
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		Name     string
	}
}

// AppConfig is the configs for the whole application
var AppConfig *Config

//Init - initialize config
func Init() error {
	if _, err := toml.DecodeFile("/home/accubits/go/src/crud/config-sample.toml", &AppConfig); err != nil {
		log.Println(" %s", err)
		return err
	}

	return nil
}
