
FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /bin/bucketlister .



FROM alpine:3.21.3

COPY --from=build /bin/bucketlister /bin/bucketlister

CMD ["/bin/bucketlister"]
