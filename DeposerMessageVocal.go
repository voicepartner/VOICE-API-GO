package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Message struct {
	ApiKey               string  `json:"apiKey"`
	TokenAudio           string  `json:"tokenAudio"`
	EmailForNotification string  `json:"emailForNotification"`
	PhoneNumbers         string  `json:"phoneNumbers"`
	ScheduledDate        *string `json:"scheduledDate,omitempty"` // Optional, use pointer to omit if nil
	Sender               *string `json:"sender,omitempty"`        // Optional, use pointer to omit if nil
	NotifyUrl            *string `json:"notifyUrl,omitempty"`     // Optional, use pointer to omit if nil
	// Additional fields can be added here
}

func main() {
	url := "https://api.voicepartner.fr/v1/campaign/send"
	//scheduledDate := "2024-04-12 14:30:00"       // Example Date and Time
	//sender := "VotreNumero"                      // Example sender number
	//notifyUrl := "https://yourdomain.com/notify" // Example notification URL

	data := Message{
		ApiKey:               "YOUR_API_KEY",
		TokenAudio:           "TOKEN_AUDIO",
		EmailForNotification: "email@exemple.com",
		PhoneNumbers:         "06xxxxxxxx",
		//ScheduledDate:        &scheduledDate, // Uncomment to use
		//Sender:               &sender,        // Uncomment to use
		//NotifyUrl:            &notifyUrl,     // Uncomment to use
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
