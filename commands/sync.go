package commands

import (
	"git.sr.ht/~porcellis/t/config"
	"os"
	"os/exec"
)

func Sync(config config.TConfig) error {
	cmd := exec.Command("git", "-C", config.BasePath, "push", config.GitRemote, "-u", "--quiet")
	cmd.Stdin = nil
	cmd.Stdout = os.NewFile(0, os.DevNull)
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
