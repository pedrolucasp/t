package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "os/exec"
  "time"
)

func generateNoteFileTitle() string {
  currentTime := time.Now()

  return currentTime.Format("2006-01-02T15:04:05Z07:00")
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

func captureText() ([]byte, error) {
  var err error
  var fileName string 

  fileName, creationErr := createNoteFile()
  
  if creationErr != nil {
    return []byte{}, creationErr
  }

  if err = executeEditor(fileName); err != nil {
    return []byte{}, err
  }
  
  bytes, err := ioutil.ReadFile(fileName)

  if err != nil {
    return []byte{}, err
  }

  fmt.Println(bytes)

  return bytes, nil
}

func notesDirectory() string {
  homeDir := os.Getenv("HOME")

  if homeDir == "" {
    fmt.Println("ERROR: You must have set your $HOME directory")
  }

  notesDirectory := fmt.Sprintf("%s/notes/", homeDir)

  return notesDirectory
}

func main() {
  fmt.Println("Starting...")
  
  data, err := captureText()

  fmt.Println(data)
  fmt.Println(err)

  fmt.Println("Finished...")
} 
