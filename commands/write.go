package commands

import (
	"git.sr.ht/~porcellis/t/models"
	"os"
	"os/exec"
)

func Write(note models.Note) error {
	cmd := exec.Command("vim", note.Path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
