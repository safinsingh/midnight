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
	"io/ioutil"
)

func dockerGen() {
	fmt.Println("Generating Dockerfile...")
	dockerfile := []byte(`FROM golang:alpine

WORKDIR /go/src/github.com/safinsingh/midnight

RUN apk add --update git
RUN go get github.com/fatih/color

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

ENTRYPOINT [ "/go/src/github.com/safinsingh/midnight/app" ]

CMD [ "pwd" ]`)

	err := ioutil.WriteFile("Dockerfile", dockerfile, 0777)

	if err != nil {
		panic(err)
	}

	fmt.Println("Dockerfile successfully generated!")
}
