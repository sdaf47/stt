package app

import (
	"github.com/jinzhu/configor"
)

const ConfigPath = "config.yml"

type Config struct {
	APPName string `default:"gid"`
}

func (app *Config) Load() (err error) {
	err = configor.Load(app, ConfigPath)
	if err != nil {
		return err
	}

	return err
}
