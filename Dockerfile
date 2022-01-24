# Start from the latest golang base image
FROM golang

RUN apt-get update --fix-missing && \
    apt-get upgrade -y && \
    apt-get install -y netcat --fix-missing


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

ENTRYPOINT ["bash", "docker-entrypoint.sh"]
