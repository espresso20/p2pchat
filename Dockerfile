FROM golang:1.17-alpine

WORKDIR /app

# Copy all files
COPY . .
COPY cert.pem /app
COPY key.pem /app

# Run the setup script
RUN ./setup.sh


# # Install gcc
# RUN apk add --no-cache gcc musl-dev

# # Run the tests
# RUN go test

# Build the server binary
RUN go build -o chatserver server.go

EXPOSE 8080

CMD ["./chatserver"]
