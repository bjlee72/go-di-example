FROM golang:1.19 as builder

# Set destination for COPY
WORKDIR /go/src 

COPY . .

# Download Go modules
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go install -v ./...

FROM ubuntu:20.04

COPY --from=builder /go/bin/di /bin/di

RUN apk add bash sudo shadow

# Run
CMD ["/bin/di"]
