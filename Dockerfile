FROM golang:1.19 as builder

# Set destination for COPY
WORKDIR /go/src 

COPY . .

# Download Go modules
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go install -v ./...

FROM scratch 
COPY --from=builder /go/bin/di /bin/di

# Run
CMD ["/bin/di"]
