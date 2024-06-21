FROM golang:1.19.11-alpine AS build
WORKDIR $GOPATH/src/crypto
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build ./cmd/main.go

FROM golang:1.19.11-alpine
WORKDIR app
COPY ./.env ./.env
COPY ./static ./static
COPY --from=build $GOPATH/src/crypto/main ./
CMD "./main"