package config

import (
	"os/user"
	"path"
)

type TConfig struct {
	BasePath  string
	GitRemote string
}

func LoadBasicConfiguration() (*TConfig, error) {
	var (
		err error
	)

	currentUser, err := user.Current()

	if err != nil {
		panic("Could not get the current user")
	}

	base := path.Join(currentUser.HomeDir, "notes")

	config := &TConfig{BasePath: base, GitRemote: "origin"}

	return config, err
}

func Initialize() (*TConfig, error) {
	return LoadBasicConfiguration()
}
