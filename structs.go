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

import "encoding/json"

// Def is the main struct for vuln definitions
type Def struct {
	OS          string  `json:"os"`
	Version     int     `json:"version"`
	Description string  `json:"description"`
	Control     Control `json:"control"`
}

// ControlType specifies the type of control of a Def
type ControlType string

// Controls are mapped to their respective types
const (
	FileControl    ControlType = "file"
	CommandControl ControlType = "command"
	PackageControl ControlType = "package"
)

// Control is the verarching struct for all control types
type Control struct {
	Type  ControlType     `json:"type"`
	Check json.RawMessage `json:"check"`
}

// FileCheck returns nil if the type is not FileControl
func (c Control) FileCheck() *FileCheck {
	if c.Type != FileControl {
		return nil
	}
	var fc FileCheck
	json.Unmarshal(c.Check, &fc)
	return &fc
}

// CommandCheck returns nil if the type is not CommandControl
func (c Control) CommandCheck() *CommandCheck {
	if c.Type != CommandControl {
		return nil
	}
	var cc CommandCheck
	json.Unmarshal(c.Check, &cc)
	return &cc
}

// PackageCheck returns nil if the type is not PackageControl
func (c Control) PackageCheck() *PackageCheck {
	if c.Type != PackageControl {
		return nil
	}
	var pc PackageCheck
	json.Unmarshal(c.Check, &pc)
	return &pc
}

// FileCheck contains the keys for a FileControl
type FileCheck struct {
	File  string `json:"file"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// CommandCheck contains the keys for a CommandControl
type CommandCheck struct {
	Command          string `json:"cmd"`
	CommandCheckType string `json:"cmdCheckType"`
	ToCheck          string `json:"toCheck"`
}

// PackageCheck contains the keys for a PackageControl
type PackageCheck struct {
	Package   string `json:"package"`
	Installed bool   `json:"installed"`
}
