FROM golang:1.18 as builder


ARG LD_FLAGS="-X github.com/sekou-diarra/fizz-buzz-server/version.Version=1.0"


WORKDIR /project
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags "$LD_FLAGS" -o fizzbuzz ./cmd/fizzbuzz

FROM alpine:3.15.1

WORKDIR /root/
COPY --from=builder /project/fizzbuzz fizzbuzz
CMD ./fizzbuzz
