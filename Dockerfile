# Pull from Alpine Docker image
FROM golang:alpine

# Set working directory to correct directory
WORKDIR /go/src/github.com/safinsingh/midnight

# Update Git within Alpine
RUN apk add --update git

# Grab dependencies
RUN go get github.com/fatih/color

# Copy files from dev environment to docker image
COPY . .

# Build executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

# Set image entrypoint to executable
ENTRYPOINT [ "/go/src/github.com/safinsingh/midnight/app" ]

# Set default flag
CMD [ "-h" ]