package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetRandomQuote() (string, error) {
	response, err := http.Get("https://zenquotes.io/api/random")

	if err != nil {
		return "", err
	}

	type quote struct {
		Content string `json:"q"`
		Author  string `json:"a"`
	}

	jsonData, _ := ioutil.ReadAll(response.Body)

	parsedData := make([]quote, 1)

	json.Unmarshal(jsonData, &parsedData)

	return fmt.Sprintf("%s - %s", parsedData[0].Content, parsedData[0].Author), nil
}
