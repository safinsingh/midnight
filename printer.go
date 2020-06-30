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
