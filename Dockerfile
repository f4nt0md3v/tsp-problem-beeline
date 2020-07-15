FROM golang:1.14-alpine AS builder

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build -o main .

# Next stage
# Use alpine or ubuntu for development
# Use scratch for production
FROM alpine

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

WORKDIR /app
COPY --from=builder /go/src/app /app

EXPOSE 9090
CMD ["./main"]
