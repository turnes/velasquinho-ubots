FROM golang:alpine as builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build  -v ./...

FROM scratch

WORKDIR /app
COPY --from=builder /go/src/app/velasquinho-ubots .
EXPOSE 5000
CMD ["./velasquinho-ubots"]