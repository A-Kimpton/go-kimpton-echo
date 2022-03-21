# Build the server binary (with CGO disabled - not compatible with docker images)
CGO_ENABLED=0 go build -o server-bin ./server

# Build the client binary (with CGO disabled - not compatible with docker images)
CGO_ENABLED=0 go build -o client-bin ./client

# Build the web binary (with CGO disabled - not compatible with docker images)
CGO_ENABLED=0 go build -o web-bin ./web

# build container for server
sudo docker build -t kimpton/go-echoserver .

# build container for client
sudo docker build -f dockerfile-client -t kimpton/go-echoclient .

# build container for web
sudo docker build -f dockerfile-web -t kimpton/go-echoweb .

# Push container to dockerhub for server
sudo docker push kimpton/go-echoserver:latest

# Push container to dockerhub for client
sudo docker push kimpton/go-echoclient:latest

# Push container to dockerhub for web
sudo docker push kimpton/go-echoweb:latest

# Run the server container
sudo docker run -p 3333:80 kimpton/go-echoweb