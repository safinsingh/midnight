#!/bin/bash

# Force script to be run as root
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit 1
fi

# Update package list
apt update

# Remove previous docker installation
apt remove docker docker-engine docker.io containerd runc

# Allow apt to use repository over HTTPS
apt install \
  apt-transport-https \
  ca-certificates \
  curl \
  gnupg-agent \
  software-properties-common

# Add GPG key for Docker installer
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -

# Add stable repository to package source list
add-apt-repository \
  "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) \
  stable"

# Update package source list
apt update

# Install Docker
apt install docker-ce docker-ce-cli containerd.io
