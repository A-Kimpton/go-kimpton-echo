# Build the server binary (with CGO disabled - not compatible with docker images)
CGO_ENABLED=0 go build -o server-bin ./server

# Build the client binary (with CGO disabled - not compatible with docker images)
CGO_ENABLED=0 go build -o client-bin ./client

# build container for server
sudo docker build -t kimpton/go-echoserver .

# build container for client
sudo docker build -f dockerfile-client -t kimpton/go-echoclient .

# Push container to dockerhub for server
#sudo docker push kimpton/go-echoserver:latest

# Push container to dockerhub for client
#sudo docker push kimpton/go-echoclient:latest

# Run the server container
sudo docker run -p 3333:9000 kimpton/go-echoclient