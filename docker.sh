#!/bin/bash

# midnight: An extensible security auditing tool
# Copyright (C) 2020  Safin Singh

# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.

# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.

# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
