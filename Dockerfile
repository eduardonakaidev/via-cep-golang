FROM golang:1.23 as builder
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o /main cmd/server/main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /main .
RUN chmod +x main
EXPOSE 3000

CMD ["./main"]