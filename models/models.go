package models

import (
	"strings"
)

type Note struct {
	Name string
	Path string
}

func (note *Note) Title() string {
	return strings.Split(note.Name, ".")[0]
}
