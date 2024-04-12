package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SmsData struct {
	ApiKey        string   `json:"apiKey"`
	Text          string   `json:"text"`
	PhoneNumbers  string   `json:"phoneNumbers"`
	Lang          *string  `json:"lang,omitempty"`          // Optional: Language of the SMS
	SpeechRate    *float64 `json:"speechRate,omitempty"`    // Optional: Rate of the speech
	NotifyUrl     *string  `json:"notifyUrl,omitempty"`     // Optional: Notification URL
	ScheduledDate *string  `json:"scheduledDate,omitempty"` // Optional: Scheduled date and time
}

func main() {
	url := "http://api.voicepartner.fr/v1/tts/send"
	//lang := "fr"
	//speechRate := 1.0
	//notifyUrl := "https://yourdomain.com/notify"
	//scheduledDate := "2024-04-12 14:30:00" // Example format: 'yyyy-MM-dd HH:mm:ss'

	data := SmsData{
		ApiKey:       "YOUR_API_KEY",
		Text:         "Votre texte ici",
		PhoneNumbers: "06XXXXXXXX",
		//Lang:         &lang,                  // Uncomment to use
		//SpeechRate:   &speechRate,            // Uncomment to use
		//NotifyUrl:    &notifyUrl,             // Uncomment to use
		//ScheduledDate: &scheduledDate,        // Uncomment to use
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
