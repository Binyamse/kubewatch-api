FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o kubewatch-api cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/kubewatch-api .
CMD ["./kubewatch-api"]