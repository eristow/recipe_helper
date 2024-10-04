# Use the official golang 1.23.2 image as the base image
FROM golang:1.23.2

# Set Ollama host env var
# ENV OLLAMA_HOST="http://localhost:11434"

# Set the working directory inside the container
WORKDIR /backend

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the Go app
RUN go build -o main ./cmd/service/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
