FROM golang:latest as builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux

ADD go.mod ./
ADD go.sum ./
RUN go mod download

ADD . .

RUN go build -ldflags="-s -w" -o goclip main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/goclip /app/goclip

EXPOSE 8080
ENTRYPOINT ["./goclip"]