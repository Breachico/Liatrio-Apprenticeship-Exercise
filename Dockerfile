# ref: https://betterstack.com/community/guides/scaling-go/dockerize-golang/
# Base image
FROM golang:1.25.3-alpine AS base

# CReate and move to working directory /app
RUN mkdir /app
WORKDIR /app

# Copy the go.mod and go.sumfiles to the build directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy source code into container
COPY . .

# Bulid application
RUN go build -o go-hello

# Inform Docker of the listening port
EXPOSE 80

# Begin application
CMD ["/app/go-hello"]
