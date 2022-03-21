#!/bin/sh

# Builds the linux binaries

# Setup env vars
export GOARCH=amd64

# Service name
service="echoservice"
rootdir=`pwd`

echo "$service :: autobuild.sh"
echo "GOBIN=$GOBIN"

# Build the client binary
echo
echo "Building the client binary $service-client"
cd $rootdir && cd client && go install $service-client.go
echo "Client complete!"

# Build the server binary
echo
echo "Building the server binary $service-server"
cd $rootdir && cd server && go install $service-server.go
echo "Server complete!"

# Build the web binary
echo
echo "Building the server binary $service-webserver"
cd $rootdir && cd web && go install $service-webserver.go
echo "Server complete!"

# Print finished message
echo 
echo "finished."