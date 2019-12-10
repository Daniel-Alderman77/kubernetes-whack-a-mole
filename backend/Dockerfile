FROM golang:latest as builder

WORKDIR /go/src/github.com/Daniel-Alderman77/kubernetes-whack-a-mole/backend

# Get dependencies
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .

# Cache downloaded dependencies
RUN go mod download

COPY . .
# Compile Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o main .

# Production stage
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy binary from previous stage
COPY --from=builder /go/src/github.com/Daniel-Alderman77/kubernetes-whack-a-mole/backend .
EXPOSE 8081
CMD ["./main"]