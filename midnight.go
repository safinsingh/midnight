package main

import "fmt"

func main() {
	config := jsonParse()

	for i := 0; i < len(config); i++ {
		fmt.Printf("[!] Check for OS %s Version %d: %s\n", config[i].Os, config[i].Version, config[i].Description)
	}
}
