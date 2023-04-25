package main

import (
	"encoding/json"
	"io"
)

var (
	catchList Catches
)

type Catch struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Gold        Gold   `json:"Gold"`
}

type Catches struct {
	Catches []Catch `json:"catches"`
}

func loadCatchAsset() {
	catchFile, err := assets.Open("assets/catch/catch.json")
	if err != nil {
		panic(err)
	}

	defer catchFile.Close()
	catches, _ := io.ReadAll(catchFile)

	if err := json.Unmarshal(catches, &catchList); err != nil {
		panic(err)
	}
}
