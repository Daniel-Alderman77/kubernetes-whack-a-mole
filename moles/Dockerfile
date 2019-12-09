FROM golang:latest as builder

WORKDIR /go/src/github.com/Daniel-Alderman77/kubernetes-whack-a-mole/moles

COPY . .

# Compile Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o main .

# Create smaller lightweight container
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy binary from previous stage
COPY --from=builder /go/src/github.com/Daniel-Alderman77/kubernetes-whack-a-mole/moles .

CMD ["./main"]