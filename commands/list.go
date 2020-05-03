package commands

import (
	"fmt"
	"git.sr.ht/~porcellis/t/config"
	"git.sr.ht/~porcellis/t/models"
	. "github.com/logrusorgru/aurora"
	"io/ioutil"
	"os"
	"path"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func List(config config.TConfig) error {
	var (
		err error
	)

	notes, err := BuildList(config)

	if err != nil {
		return err
	}

	fmt.Println(Bold("\n \t Your notes\n\n"))

	for index, note := range notes {
		fmt.Printf("[#%d] %s (%s)\n", index, Bold(note.Title()), Faint(note.UpdatedAt()).BrightYellow())

		f, err := os.Open(note.Path)
		check(err)

		b1 := make([]byte, 35)
		n1, err := f.Read(b1)
		check(err)
		fmt.Printf("%s\n\n", string(b1[:n1]))
	}

	return err
}

func BuildList(config config.TConfig) ([]models.Note, error) {
	var (
		err   error
		notes []models.Note
	)

	files, err := ioutil.ReadDir(config.BasePath)

	sort.Slice(files, func(index, aux int) bool {
		return files[index].ModTime().After(files[aux].ModTime())
	})

	for _, entry := range files {
		if !entry.IsDir() {
			note := models.Note{Name: entry.Name(), Path: path.Join(config.BasePath, entry.Name()), ModTime: entry.ModTime()}

			notes = append(notes, note)
		}
	}

	return notes, err
}
