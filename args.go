package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func banner(mode string, defs string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println(`
███╗   ███╗██╗██████╗ ███╗   ██╗██╗ ██████╗ ██╗  ██╗████████╗
████╗ ████║██║██╔══██╗████╗  ██║██║██╔════╝ ██║  ██║╚══██╔══╝
██╔████╔██║██║██║  ██║██╔██╗ ██║██║██║  ███╗███████║   ██║   
██║╚██╔╝██║██║██║  ██║██║╚██╗██║██║██║   ██║██╔══██║   ██║   
██║ ╚═╝ ██║██║██████╔╝██║ ╚████║██║╚██████╔╝██║  ██║   ██║   
╚═╝     ╚═╝╚═╝╚═════╝ ╚═╝  ╚═══╝╚═╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝   `)
	fmt.Printf("Running in ")
	infoPrint(mode)
	fmt.Printf(" mode\n")
	fmt.Printf("Defs: ")
	infoPrint(defs)

	fmt.Println()
	fmt.Println()
	fmt.Print("Press 'Enter' to proceed...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	fmt.Println()
}

func argparse() (string, string) {
	mode := flag.String("mode", "audit", "Mode to run midnight in. Possible modes: audit, enforce, docker")
	file := flag.String("file", "", "Configuration file to use (mandatory)")

	flag.Parse()

	if *file == "" {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	banner(*mode, *file)

	return *mode, *file
}
