package main

import (
	"git.sr.ht/~porcellis/t/commands"
	"git.sr.ht/~porcellis/t/config"
	"git.sr.ht/~porcellis/t/models"
	"git.sr.ht/~sircmpwn/getopt"
	"os"
)

func main() {
	var (
		c *config.TConfig
	)

	c, err := config.Initialize()

	if err != nil {
		panic(err)
	}

	opts, optind, err := getopt.Getopts(os.Args, "le:c:d:")

	if err != nil {
		panic(err)
	}

	for _, opt := range opts {
		switch opt.Option {
		case 'e':
			println("Edit specified")
		case 'l':
			commands.List(*c)
		case 'c':
			println(opt.Value)

			var note models.Note
			if opt.Value != "" {
				note, err = commands.CreateWithName(*c, opt.Value)
			} else {
				note, err = commands.Create(*c)
			}

			if err != nil {
				panic("We we're not able to create your note")
			}

			err = commands.Write(note)

			if err != nil {
				panic("We could not invoke your $EDITOR")
			}

			// We should call commit
			// commands.Commit()

			// commands.Sync()
			println("Note created", note.Name)
		case 'd':
			println("Delete specified")
		}
	}

	for _, arg := range os.Args[optind:] {
		if string(arg) == "l" {
			commands.List(*c)
		}
	}

}
