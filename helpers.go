package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func GetRandomJoke(category string) (string, error) {

	validCategories := []string{"Any", "Misc", "Programming", "Dark", "Pun", "Spooky", "Christmas"}

	var valid bool = false

	for _, validCategory := range validCategories {
		if strings.EqualFold(category, validCategory) {
			valid = true
			category = validCategory
		}
	}

	if category == "" {
		category = "Any"
		valid = true
	}

	if !valid {
		return "", errors.New("invalid joke category")
	}

	response, err := http.Get("https://v2.jokeapi.dev/joke/" + category)

	if err != nil {
		return "", err
	}

	type joke struct {
		Type     string `json:"type"`
		Setup    string `json:"setup"`
		Delivery string `json:"delivery"`
		Joke     string `json:"joke"`
	}

	jsonData, _ := ioutil.ReadAll(response.Body)

	parsedData := &joke{}

	err = json.Unmarshal(jsonData, &parsedData)

	if err != nil {
		return "", nil
	}

	var result = parsedData.Joke

	if parsedData.Type == "twopart" {
		result = parsedData.Setup + "\n" + parsedData.Delivery
	}

	return result, nil
}
