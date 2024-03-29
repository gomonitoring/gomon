# Start from the latest golang base image
FROM golang:alpine as builder

ENV FUNC=server

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
COPY . .
RUN go build -o main .

FROM alpine
# Copy the source from the current directory to the Working Directory inside the container
COPY --from=builder /app/main /app/docker-entrypoint.sh ./

RUN apk add bash

ENTRYPOINT ["bash", "docker-entrypoint.sh"]
