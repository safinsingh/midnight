package main

import (
	"encoding/json"
	"io/ioutil"
)

func jsonParse() []Def {
	data, err := ioutil.ReadFile("checks/u16stig.json")
	if err != nil {
		panic(err)
	}

	var defs []Def
	if err := json.Unmarshal([]byte(data), &defs); err != nil {
		panic(err)
	}

	return defs
}
