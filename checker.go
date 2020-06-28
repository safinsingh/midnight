package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

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

func check(control Def) bool {
	if !keyInFile(control.Control.Key, control.Control.Location) {
		return false
	}

	if keyPairInFile(control.Control.Key, control.Control.Value, control.Control.Location) {
		return true
	}

	return false
}

func commence(controls []Def) {
	for i := 0; i < len(controls); i++ {
		code := check(controls[i])
		if code {
			successPrint(controls[i])
		} else {
			failPrint(controls[i])
		}
	}
}
