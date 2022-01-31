# Start from the latest golang base image
FROM golang:alpine as builder

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    apk add netcat-openbsd


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN go build -o main 

FROM alpine:3.11.3
# Copy the source from the current directory to the Working Directory inside the container
COPY --from=builder . .

ENTRYPOINT ["bash", "docker-entrypoint.sh"]
