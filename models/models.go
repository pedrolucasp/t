package models

import (
	"strings"
	"time"
)

type Note struct {
	Name    string
	Path    string
	ModTime time.Time
}

func (note *Note) Title() string {
	return strings.Split(note.Name, ".")[0]
}

func (note *Note) UpdatedAt() string {
	return note.ModTime.Format("Mon, Jan _2 15:04:05 2006")
}
