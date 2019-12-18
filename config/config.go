package config

import (
  "path"
  "fmt"
  "os/user"
)

type TConfig struct {
  BasePath string
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

  fmt.Println(base)

  config := &TConfig{BasePath: base, GitRemote: "origin"} 

  return config, err
}

func Initialize() (*TConfig, error) {
  return LoadBasicConfiguration()
}
