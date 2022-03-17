# Build the server binary (with CGO disabled - not compatible with docker images)
CGO_ENABLED=0 go build -o main ./server

# build container
sudo docker build -t kimpton/go-echoserver .

# Push container to dockerhub
sudo docker push kimpton/go-echoserver:latest

# Run the container
sudo docker run -p 3333:9000 go-echoserver