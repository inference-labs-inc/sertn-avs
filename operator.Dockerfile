FROM golang:1.21 as build

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod tidy && go mod verify

COPY . .

WORKDIR /usr/src/app/operator/cmd
RUN go build -v -o /usr/local/bin/operator ./...

FROM debian:latest
COPY --from=build /usr/local/bin/operator /usr/local/bin/operator
ENTRYPOINT [ "operator"]
CMD ["--config=/app/avs_config.yaml"]