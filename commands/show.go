package commands

import (
	"git.sr.ht/~porcellis/t/models"
	"os"
	"os/exec"
)

func Show(note models.Note) error {
	// TODO: Make this configurable
	cmd := exec.Command("glow", note.Path, "-p")
	cmd.Stdin = nil
	cmd.Stdout = os.NewFile(0, os.DevNull)
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
