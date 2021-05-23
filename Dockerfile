FROM golang:1.16.4-alpine3.13 AS build

COPY go.mod go.sum /src/
WORKDIR /src
RUN go mod download

COPY . /src
WORKDIR /src/
RUN go build -o dispatcher pkg/main.go

FROM golang:1.16.4-alpine3.13

WORKDIR /go
COPY ./configs/ /go/configs/
COPY --from=build /src/dispatcher /go/

ENV PATH="/go:${PATH}"

EXPOSE 3000
CMD ["dispatcher"]