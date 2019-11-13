package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "os/exec"
  "flag"
  "time"
)

var noteName string

func generateNoteTitle() string {
  currentTime := time.Now()

  noteName = currentTime.Format("2006-01-02T15:04:05Z07:00")
  
  return noteName
}

func noteFilePath() string {
  return fmt.Sprintf("%s%s.md", notesDirectory(), noteName)
}

func createNoteFile() error {
  file, err := os.Create(noteFilePath())
  
  if err != nil {
    fmt.Println("Could not create the file")
    fmt.Println(err)
  }

  file.Close()

  return err
}

func executeEditor() error {
  cmd := exec.Command("vim", noteFilePath())
  cmd.Stdin  = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  return cmd.Run()
}

func captureText() error {
  var err error

  fmt.Println("Capturing...")

  if err = executeEditor(); err != nil {
    return err
  }
  
  bytes, err := ioutil.ReadFile(noteFilePath())

  if err != nil {
    return err
  }

  fmt.Println(bytes)

  return nil
}

func isAGitRepository() error {
  if _, err := os.Stat(fmt.Sprintf("%s.git", notesDirectory())); os.IsNotExist(err) {
	  return err
  }

  return nil
}

func notesDirectory() string {
  homeDir := os.Getenv("HOME")

  if homeDir == "" {
    fmt.Println("ERROR: You must have set your $HOME directory")
  }

  notesDirectory := fmt.Sprintf("%s/notes/", homeDir)

  return notesDirectory
}

func syncNotes() error {
  cmd := exec.Command("git", "-C", notesDirectory(), "push", "--set-upstream", "origin", "master")
  cmd.Env = append(os.Environ())

  cmd.Stdin  = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  err := cmd.Run()
  
  if err == nil {
    fmt.Println("It worked")
    return nil
  } else {
    return err
  }
}

func addChangedFiles() error {
  cmd := exec.Command("git", "-C", notesDirectory(), "add", "-A")
  cmd.Stderr = os.Stderr
  
  err := cmd.Run()

  return err
}

func commitLastNote() error {
  // TODO: We should pretty format this date
  // Also handle different actions as well

  addingError := addChangedFiles()
  
  if addingError != nil {
    return addingError
  }

  commitMessage := fmt.Sprintf("Added note '%s'", noteName)

  cmd := exec.Command("git", "-C", notesDirectory(), "commit", "-am", commitMessage)
  cmd.Env = append(os.Environ())
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  err := cmd.Run()

  if err != nil {
    return err
  } else {
    return nil
  }
}

func create(name string) error {
  noteName = name

  if err := createNoteFile(); err == nil {
    if err := captureText(); err != nil {
      fmt.Println("Could not write to the note")
      return err
    } else {
      sync()

      return nil
    }  
  } else {
    fmt.Println("Could not create the note")
    return err
  }
}

func edit(index int) {
  fmt.Println("Edit...")
}

func sync() {
  fmt.Println("Sync")
  if err := commitLastNote(); err == nil {
    if err := syncNotes(); err == nil {
      fmt.Println("Finished sync :)")
    } else {
      fmt.Println("Something wrong happened")
    }  
  } else {
    fmt.Println("Could not commit the note")
  }
}

func main() {
  addCommand := flag.NewFlagSet("add", flag.ExitOnError)
  addCommandName := addCommand.String("name", "", "name")

  editCommand := flag.NewFlagSet("edit", flag.ExitOnError)
  editCommandIndex := editCommand.Int("index", 0, "index")

  if len(os.Args) > 1 {
    switch os.Args[1] {
      case "add":
        addCommand.Parse(os.Args[2:])
        if *addCommandName == "" {
          create(generateNoteTitle())
        } else {
          create(*addCommandName)
        }
      case "edit":
        editCommand.Parse(os.Args[2:])
        fmt.Println("index: ", *editCommandIndex)
      default:
        fmt.Println("Will start add process")
        create(generateNoteTitle())
    }
  } else {
    fmt.Println("Start add process")
    create(generateNoteTitle())
  }
  
}
