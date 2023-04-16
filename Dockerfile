FROM golang:1.19.3-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main app/cmd/main/main.go


FROM alpine

WORKDIR /app

COPY --from=builder main .

CMD ["/app/main"]