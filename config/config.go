package config

import (
	"log"
	"path"

	"github.com/go-ini/ini"
	"github.com/kyoh86/xdg"
)

type ShareConfig struct {
	Protocol string `ini:"protocol"`
	Base     string `ini:"base"`
	Path     string `ini:"path"`
}

type SyncConfig struct {
	Remote string `ini:"remote"`
}

type TConfig struct {
	BasePath string      `ini:"base"`
	Share    ShareConfig `ini:"-"`
	Sync     SyncConfig  `ini:"-"`
}

func LoadConfigFromFile(root *string, sharedir string) (*TConfig, error) {
	// TODO: Handle loading from the default config
	filename := path.Join(xdg.ConfigHome(), "t", "t.conf")

	file, err := ini.LoadSources(ini.LoadOptions{
		KeyValueDelimiters: "=",
	}, filename)

	if err != nil {
		// TODO: Install if was not installed within the make command
		log.Fatalln("Can't find config")
		return nil, err
	}

	// TODO: Handle converting ~/ into the user default home directory
	config := &TConfig{
		Share: ShareConfig{
			Protocol: "https",
			Base:     "rascunho.eletrotupi.com",
			Path:     "/api/v1",
		},

		Sync: SyncConfig{
			Remote: "origin",
		},
	}

	err = file.MapTo(&config)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func LoadBasicConfiguration() (*TConfig, error) {
	var (
		err error
	)

	config, err := LoadConfigFromFile(nil, "")

	if err != nil {
		return nil, err
	}

	return config, err
}

func Initialize() (*TConfig, error) {
	return LoadBasicConfiguration()
}
