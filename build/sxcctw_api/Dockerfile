# First stage: Build the Go application
FROM golang:alpine AS build

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy the rest of the application files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app ./cmd/sxcctw_api/

# Second stage: Serve the Go application
FROM alpine

# Install package for timezone.
RUN apk update && apk --no-cache add tzdata

# Copy the built binary from the build stage
COPY --from=build /go/bin/app /go/bin/app

# Start the Go application
CMD ["/go/bin/app"]
