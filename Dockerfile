# Use the official Golang image as the base image
FROM golang:1.22.5-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

ARG DB_USER
ARG DB_PASS
ARG DB_HOST
ARG DB_PORT
ARG DB_NAME

# Create a .env file with the environment variables
RUN echo "DB_USER=${DB_USER}" > .env && \
    echo "DB_PASS=${DB_PASS}" >> .env && \
    echo "DB_HOST=${DB_HOST}" >> .env && \
    echo "DB_PORT=${DB_PORT}" >> .env && \
    echo "DB_NAME=${DB_NAME}" >> .env

RUN ls -la

RUN cat .env

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 80

# Command to run the executable
CMD ["./main"]
