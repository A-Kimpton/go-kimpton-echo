# Define base image - maybe alpine is not a good fit
FROM golang:alpine

# Make a working dir
RUN mkdir /es-server

# Copy the src to the working dir
COPY server-bin /es-server/

# Set the working dir
# WORKDIR /es-server

# Set permission to allow execute
RUN chmod +x /es-server/server-bin

RUN echo $(ls -l /es-server/)

# Run the app
CMD ["/es-server/server-bin"]