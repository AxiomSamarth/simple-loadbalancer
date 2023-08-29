# Use a suitable base image, for example, one that supports Go applications
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the project
RUN go build -o app main.go

# Expose port 8080
EXPOSE 8080

# Command to run your application
CMD ["./app"]
