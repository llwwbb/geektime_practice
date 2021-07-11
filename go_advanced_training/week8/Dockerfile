FROM golang:1.16-alpine3.14 as builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -o main

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main ./
CMD ["./main"]