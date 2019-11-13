package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "os/exec"
  "time"
)

var noteName string

func generateNoteFileTitle() string {
  currentTime := time.Now()

  noteName = currentTime.Format("2006-01-02T15:04:05Z07:00")
  
  return noteName
}

func createNoteFile() (string, error) {
  file, err := os.Create(fmt.Sprintf("%s%s.md", notesDirectory(), generateNoteFileTitle()))

  if err != nil {
    fmt.Println("Could not create the file")
    fmt.Println(err)
  }

  fileName := file.Name()
  file.Close()

  return fileName, err
}

func executeEditor(fileName string) error {
  cmd := exec.Command("vim", fileName)
  cmd.Stdin  = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  return cmd.Run()
}

func captureText() (string, error) {
  var err error
  var fileName string 

  fileName, creationErr := createNoteFile()
  
  if creationErr != nil {
    return "", creationErr
  }

  if err = executeEditor(fileName); err != nil {
    return "", err
  }
  
  bytes, err := ioutil.ReadFile(fileName)

  if err != nil {
    return "", err
  }

  fmt.Println(bytes)

  return fileName, nil
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

func commitLastNote(fileName string) error {
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

func main() {
  fmt.Println("Starting...")
  
  lastNote, err := captureText()

  fmt.Println(lastNote)
  
  if err != nil {
    panic(err)
  } else {
    commitLastNote(lastNote)
  }

  fmt.Println("Syncing your notes...")

  if err := isAGitRepository(); err != nil {
    fmt.Println("This is not a valid repository")
  } else {
    if err := syncNotes(); err != nil {
      fmt.Println("Could not sync notes")
      fmt.Println(err)
    }
  }

  fmt.Println("Finished...")
}
