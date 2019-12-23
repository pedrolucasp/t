package commands

import (
	"git.sr.ht/~porcellis/t/models"
	"os"
	"os/exec"
)

func Write(note models.Note) error {
	var (
		editor string
	)

	editor = os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	cmd := exec.Command(editor, note.Path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
