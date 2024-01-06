FROM golang:1.19

# Set destination for COPY
WORKDIR /go/src

COPY . .

# Download Go modules
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go install -v ./...

# Run
CMD ["/go/bin/di"]
