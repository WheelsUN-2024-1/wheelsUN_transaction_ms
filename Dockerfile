# Use the official Go image as a parent image
FROM golang:1.18 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Fetch dependencies.
# Using go mod with go 1.11 or later
RUN go mod download
RUN go mod verify

# Build the command inside the container.
# (It may be preferable to build using the -ldflags "-s" to strip the debugging information)
RUN CGO_ENABLED=0 GOOS=linux go build -v -o myapp

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/
# Use the official Alpine image for a lean production container.
# Only copy the binary from builder stage
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/myapp .

# Run the binary program produced by go install
CMD ["./myapp"]