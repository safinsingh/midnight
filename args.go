// midnight: An extensible security auditing tool
// Copyright (C) 2020  Safin Singh

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
