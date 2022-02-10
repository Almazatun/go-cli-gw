package main

import (
	"encoding/json"
	"io/ioutil"
)

type Pass struct {
	Username    string
	Email       string
	Password    string
	Platform    string
	Description string
}

func getPass() (pass []Pass) {
	fileBytes, err := ioutil.ReadFile("./pass-list.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &pass)

	if err != nil {
		panic(err)
	}

	return pass
}

func savePass(passList []Pass) {
	passBytes, err := json.Marshal(passList)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./pass-list.json", passBytes, 0644)

	if err != nil {
		panic(err)
	}
}
