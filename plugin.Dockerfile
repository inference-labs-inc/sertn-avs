FROM golang:1.21 as build

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod tidy && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/plugin plugin/cmd/main.go

FROM debian:latest
COPY --from=build /usr/local/bin/plugin /usr/local/bin/plugin
ENTRYPOINT [ "plugin", "--config=/app/avs_config.yaml" ]