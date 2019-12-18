package main

import (
  "fmt"
  "git.sr.ht/~porcellis/t/models"
  "git.sr.ht/~porcellis/t/config"
)

func main() {
  var (
    n models.Note
    c *config.TConfig
  )

  c, err := config.Initialize()
  
  if err != nil {
    panic(err)
  }

  fmt.Println(c)

  n = models.Note{"Test", "~/notes/test.md"}

  fmt.Println(n)
}
