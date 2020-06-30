#!/bin/bash

# Update package list
apt update

# Install golang and git (for go get)
apt install -y golang-go git

# Grab dependenciy
go get "github.com/fatih/color"

# Build executable
go build .