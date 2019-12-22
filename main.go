package main

import (
	"fmt"
	"git.sr.ht/~porcellis/t/commands"
	"git.sr.ht/~porcellis/t/config"
)

func main() {
	var (
		c *config.TConfig
	)

	c, err := config.Initialize()

	if err != nil {
		panic(err)
	}

	err = commands.List(*c)

	if err != nil {
		fmt.Println(err)
	}
}
