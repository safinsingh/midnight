package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func jsonParse() []Def {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config.json")
	}

	var def []Def
	json.Unmarshal([]byte(data), &def)

	return def
}
