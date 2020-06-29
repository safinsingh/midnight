package main

import (
	"fmt"

	"github.com/fatih/color"
)

func infoPrint(val string) {
	cyan := color.New(color.FgCyan, color.Bold)
	cyan.Printf("%s", val)
}

func successPrint(control Def) {
	green := color.New(color.FgGreen, color.Bold)

	green.Printf("[+] ")
	fmt.Printf("Check for OS: ")
	green.Printf("%s %d", control.OS, control.Version)
	fmt.Printf(" - ")
	green.Printf(control.Description)
	fmt.Printf("\n")
}

func failPrint(control Def) {
	red := color.New(color.FgRed, color.Bold)

	red.Printf("[-] ")
	fmt.Printf("Check for OS: ")
	red.Printf("%s %d", control.OS, control.Version)
	fmt.Printf(" - ")
	red.Printf(control.Description)
	fmt.Printf("\n")
}
