package main

import (
  "fmt"
  "git.sr.ht/~porcellis/t/models"
)

func main() {
  var (
    n models.Note
  )

  n = models.Note{"Test", "~/notes/test.md"}

  fmt.Println(n)
}
