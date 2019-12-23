package main

import (
	"git.sr.ht/~porcellis/t/commands"
	"git.sr.ht/~porcellis/t/config"
	"git.sr.ht/~porcellis/t/models"
	"git.sr.ht/~sircmpwn/getopt"
	"log"
	"os"

	"strconv"
)

var Version = "0.0.1"

func usage() {
	log.Fatal("Usage: t -[l,c,e]")
}

func main() {
	var (
		c *config.TConfig
	)

	c, err := config.Initialize()

	if err != nil {
		panic(err)
	}

	opts, _, err := getopt.Getopts(os.Args, "vle:c:d:")

	if err != nil {
		usage()

		return
	}

	for _, opt := range opts {
		switch opt.Option {
		case 'e':
			println("Editing")

			var note models.Note
			notes, _ := commands.BuildList(*c)

			if opt.Value == "" {
				note = notes[0]
			} else {
				index, err := strconv.Atoi(opt.Value)

				if err == nil {
					note = notes[index]
				}
			}

			err = commands.Write(note)

			err = commands.Commit(*c, note)

			if err != nil {
				panic("Could not commit your edited note")
			}

			err = commands.Sync(*c)

			if err != nil {
				panic("Could not sync your note")
			}

			println("Finished editing ", note.Title())
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
			err = commands.Commit(*c, note)

			if err != nil {
				panic("We could not commit your note")
			}

			err = commands.Sync(*c)

			if err != nil {
				panic("We could not sync your notes")
			}

			println("Note created", note.Name)
		case 'v':
			println("t", Version)
		}
	}
}
