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
	"encoding/json"
	"io/ioutil"
)

func jsonParse(file string) []Def {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var defs []Def
	if err := json.Unmarshal([]byte(data), &defs); err != nil {
		panic(err)
	}

	return defs
}
