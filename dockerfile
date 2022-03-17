# Define base image - maybe alpine is not a good fit
FROM golang:alpine

# Make a working dir
RUN mkdir /es-server

# Copy the src to the working dir
COPY main /es-server/

# Set the working dir
# WORKDIR /es-server

# Set permission to allow execute
RUN chmod +x /es-server/main

RUN echo $(ls -l /es-server/)

# Run the app
CMD ["/es-server/main"]