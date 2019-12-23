package commands

import (
	"fmt"
	"git.sr.ht/~porcellis/t/config"
	"git.sr.ht/~porcellis/t/models"
	"os"
	"os/exec"
)

func AddToStagedList(basePath string, notePath string) error {
	cmd := exec.Command("git", "-C", basePath, "add", notePath)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Commit(config config.TConfig, note models.Note) error {
	var (
		err error
	)

	err = AddToStagedList(config.BasePath, note.Path)

	if err != nil {
		return err
	}

	commitMessage := fmt.Sprintf("Modified '%s'", note.Title())

	cmd := exec.Command("git", "-C", config.BasePath, "commit", "-am", commitMessage)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
