package main

import (
	"fmt"
	"git.sr.ht/~porcellis/t/commands"
	"git.sr.ht/~porcellis/t/config"
	"git.sr.ht/~porcellis/t/models"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var Version = "0.1.0"

func usage() {
	log.Fatal("Usage: t [list, create, edit, version]")
}

func main() {
	var (
		configuration *config.TConfig
	)

	configuration, err := config.Initialize()

	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		usage()

		return
	}

	switch os.Args[1] {
	case "create", "c":
		var note models.Note
		if len(os.Args) == 2 {
			note, err = commands.Create(*configuration)
		} else {
			note, err = commands.CreateWithName(*configuration, os.Args[2])
		}

		if err != nil {
			panic("We we're not able to create your note")
		}

		err = commands.Write(note)

		if err != nil {
			panic("We could not invoke your $EDITOR")
		}

		// We should call commit
		err = commands.Commit(*configuration, note)

		if err != nil {
			panic("We could not commit your note")
		}

		err = commands.Sync(*configuration)

		if err != nil {
			panic("We could not sync your notes")
		}

		fmt.Println(fmt.Sprintf("Note created %s", note.Name))

	case "list", "l":
		content, err := commands.List(*configuration)

		if err != nil {
			panic("We could not fetch your notes")
		}

		pager := os.Getenv("PAGER")

		if pager == "" {
			// This options display control chars as raw
			pager = "less -r"
		}

		pagerCommand := strings.Split(pager, " ")

		cmd := exec.Command(pagerCommand[0], pagerCommand[1:]...)
		cmd.Stdin = strings.NewReader(content)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.Run()

	case "edit", "e":
		var note models.Note
		notes, _ := commands.BuildList(*configuration)

		if len(os.Args) == 2 {
			note = notes[0]
		} else {
			index, err := strconv.Atoi(os.Args[2])

			if err == nil {
				note = notes[index]
			}
		}

		err = commands.Write(note)

		err = commands.Commit(*configuration, note)

		if err != nil {
			panic("Could not commit your edited note")
		}

		err = commands.Sync(*configuration)

		if err != nil {
			panic("Could not sync your note")
		}

		println("Finished editing ", note.Title())

	case "show", "s":
		var note models.Note
		notes, _ := commands.BuildList(*configuration)

		if len(os.Args) == 2 {
			note = notes[0]
		} else {
			index, err := strconv.Atoi(os.Args[2])

			if err == nil {
				note = notes[index]
			}
		}

		err = commands.Show(note)

		if err != nil {
			panic("There was some error when trying to display the note")
		}

	case "share", "sh":
		var note models.Note
		notes, _ := commands.BuildList(*configuration)

		if len(os.Args) == 2 {
			note = notes[0]
		} else {
			index, err := strconv.Atoi(os.Args[2])

			if err == nil {
				note = notes[index]
			}
		}

		err = commands.Share(*configuration, note)

		if err != nil {
			panic(fmt.Sprintf("There was some error when trying to share the note, %s", err))
		}

	case "version", "v":
		println("t ", Version)

	case "help", "h":
		usage()

	default:
		usage()
		os.Exit(1)
	}
}
