package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.sr.ht/~porcellis/t/config"
	"git.sr.ht/~porcellis/t/models"
	"io/ioutil"
	"log"
	"net/http"
)

func Share(config config.TConfig, note models.Note) error {
	url := fmt.Sprintf("%s://%s", config.Share.Protocol, config.Share.Base)
	api := fmt.Sprintf("%s%s", url, config.Share.Path)

	contents, err := ioutil.ReadFile(note.Path)

	if err != nil {
		return err
		log.Fatalln(err)
	}

	requestBody, err := json.Marshal(map[string]string{
		"text": string(contents),
	})

	// TODO: Make this configurable
	response, err := http.Post(api, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
		log.Fatalln(err)
	}

	defer response.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

	log.Println(fmt.Sprintf("%s/%s", url, result["sha"]))

	return nil
}
