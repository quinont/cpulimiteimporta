FROM golang:1.19.2-alpine3.16 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build

COPY . .

RUN go build -o main main.go

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=builder /dist/main /

ENTRYPOINT ["/main"]
