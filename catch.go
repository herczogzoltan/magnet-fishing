package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	catchList Catches
)

type Catch struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type Catches struct {
	Catches []Catch `json:"catches"`
}

func loadCatchAsset() {
	jsonFile, err := os.Open("assets/catch/catch.json")
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()
	catches, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(catches, &catchList)
	if err != nil {
		panic(err)
	}
}
