package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiKey := "YOUR_API_KEY"
	campaignId := "CAMPAIGN_ID"
	url := fmt.Sprintf("https://api.voicepartner.fr/v1/campaign/cancel/%s/%s", apiKey, campaignId)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

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
