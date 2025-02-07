FROM golang:1.23 AS build

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o mahoor

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app .

EXPOSE 7777

ENV GIN_MODE=release

CMD ["./mahoor"]
