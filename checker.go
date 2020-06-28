package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

////// FILE CONTROL FUNCTIONS //////

func keyInFile(key string, file string) bool {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return strings.Contains(string(f), key)
}

func keyPairInFile(key string, val string, file string) bool {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), key) && strings.Contains(scanner.Text(), val) {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return false
}

func fileControl(control FileCheck) bool {
	if !keyInFile(control.Key, control.File) {
		return false
	}

	if keyPairInFile(control.Key, control.Value, control.File) {
		return true
	}

	return false
}

////// COMMAND CHECKS //////

func cmdControl(control CommandCheck) bool {
	return true
}

////// MAIN SWITCHER //////

func checkSwitch(control Def) bool {
	switch control.Control.Type {
	case "file":
		return fileControl(*control.Control.FileCheck())
	case "command":
		return cmdControl(*control.Control.CommandCheck())
	}

	return false
}

func commence(controls []Def) {
	for i := 0; i < len(controls); i++ {
		code := checkSwitch(controls[i])

		if code {
			successPrint(controls[i])
		} else {
			failPrint(controls[i])
		}
	}
}
