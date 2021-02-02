FROM golang:alpine as builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -v ./...

FROM scratch

WORKDIR /app

COPY --from=builder /go/src/app/velasquinho-ubots .

CMD ["./velasquinho-ubots"]