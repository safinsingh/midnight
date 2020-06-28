package main

import (
	"encoding/json"
	"io/ioutil"
)

func jsonParse() []Def {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var def []Def
	json.Unmarshal([]byte(data), &def)

	return def
}
