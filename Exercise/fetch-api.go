package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {
	url := "https://catfact.ninja/fact"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the API:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var catFact CatFact
	err = json.Unmarshal(body, &catFact)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Random Cat Fact:", catFact.Fact)
}
