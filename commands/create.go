package commands

import (
	"fmt"
	"git.sr.ht/~porcellis/t/config"
	"git.sr.ht/~porcellis/t/models"
	"os"
	"path"
	"time"
)

func CreateWithName(config config.TConfig, title string) (models.Note, error) {
	var (
		err  error
		note models.Note
	)

	file, err := os.Create(path.Join(config.BasePath, fmt.Sprintf("%s.md", title)))

	stat, _ := file.Stat()

	note = models.Note{Name: stat.Name(), Path: path.Join(config.BasePath, stat.Name()), ModTime: stat.ModTime()}

	return note, err

}

func Create(config config.TConfig) (models.Note, error) {
	var (
		err  error
		note models.Note
	)

	currentTime := time.Now()
	noteName := currentTime.Format("2006-01-02T15:04:05Z07:00")

	file, err := os.Create(path.Join(config.BasePath, fmt.Sprintf("%s.md", noteName)))

	stat, _ := file.Stat()
	note = models.Note{Name: stat.Name(), Path: path.Join(config.BasePath, stat.Name()), ModTime: stat.ModTime()}

	return note, err
}
