package commands

import (
	"fmt"
	"git.sr.ht/~porcellis/t/config"
	"git.sr.ht/~porcellis/t/models"
	"io/ioutil"
	"path"
	"sort"
)

func List(config config.TConfig) error {
	var (
		err error
	)

	notes, err := BuildList(config)

	if err != nil {
		return err
	}

	fmt.Println("\n \t Your notes")

	for index, note := range notes {
		if index > 0 {
			fmt.Printf("|#%-6d|\t%6s\t|%6s|\n", index, note.Title(), note.UpdatedAt())
		}
	}

	return err
}

func BuildList(config config.TConfig) ([]models.Note, error) {
	var (
		err   error
		notes []models.Note
	)

	fmt.Println("Listing some stuff")

	files, err := ioutil.ReadDir(config.BasePath)

	sort.Slice(files, func(index, aux int) bool {
		return files[index].ModTime().After(files[aux].ModTime())
	})

	for _, entry := range files {
		note := models.Note{Name: entry.Name(), Path: path.Join(config.BasePath, entry.Name()), ModTime: entry.ModTime()}

		notes = append(notes, note)
	}

	return notes, err
}
