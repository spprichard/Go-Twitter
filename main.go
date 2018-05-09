package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var keys struct {
		Key    string `json:"consumer_key"`
		Secret string `json:"consumer_secret"`
	}

	file, err := os.Open(".keys.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&keys)

	fmt.Printf("%+v\n", keys)
}
