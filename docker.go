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
