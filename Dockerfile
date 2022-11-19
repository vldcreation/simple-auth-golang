FROM golang:1.17-alpine3.15

LABEL maintainer="Vicktor Desrony"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Build the Go app
RUN go build -o /build

# Expose port 9091 to the outside world
EXPOSE 9091

# Run the executable
CMD [ "/build" ]