package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://api.voicepartner.fr/v1/audios/YOUR_API_KEY"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Response: %s\n", string(body))
}
