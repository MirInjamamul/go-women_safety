# Use the official golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code from the host machine to the container
COPY ./cmd /app/cmd

# Build the Go application inside the container
RUN go build -o myapp ./cmd/main.go

# Expose the port that the application will listen on
EXPOSE 8080

# Set the command to run when the container starts
CMD ["./myapp"]