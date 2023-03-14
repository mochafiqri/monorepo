# Use golang as the builder image
FROM golang:latest as builder

# Create the app directory and copy the source files
RUN mkdir /app
COPY . /app

# Set the working directory and build the binary
WORKDIR /app/cmd
RUN go build -o main .

# Set the entrypoint to the binary
CMD ["/app/main"]
