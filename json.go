package main

import (
	"encoding/json"
	"io/ioutil"
)

func jsonParse(file string) []Def {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var defs []Def
	if err := json.Unmarshal([]byte(data), &defs); err != nil {
		panic(err)
	}

	return defs
}
