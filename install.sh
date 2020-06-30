#!/bin/bash

# Force script to be run as root
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit 1
fi

# Update package list
apt update

# Install golang and git (for go get)
apt install -y golang-go git

# Grab dependenciy
go get "github.com/fatih/color"

# Build executable
go build .