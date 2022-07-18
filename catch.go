package main

import (
	_ "embed"
	"encoding/json"
)

var (
	catchList Catches
	//go:embed assets/catch/catch.json
	catches []byte
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
	if err := json.Unmarshal(catches, &catchList); err != nil {
		panic(err)
	}
}
