package main

import "fmt"

func handle(defs []Def, mode string) {
	switch mode {
	case "audit":
		commence(defs)
	case "enforce":
		fmt.Println("Not yet supported")
	case "docker":
		dockerGen()
	}
}
