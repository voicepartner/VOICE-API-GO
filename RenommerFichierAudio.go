package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RenameData struct {
	ApiKey     string `json:"apiKey"`
	TokenAudio string `json:"tokenAudio"`
	Filename   string `json:"filename"`
}

func main() {
	url := "https://api.voicepartner.fr/v1/audio-file/rename"
	data := RenameData{
		ApiKey:     "YOUR_API_KEY",
		TokenAudio: "TOKEN_DU_FICHIER_AUDIO",
		Filename:   "Nom du fichier",
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Response: %s\n", string(respBody))
}
