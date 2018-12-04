FROM golang:latest AS builder

WORKDIR /go/src/github.com/cjhall1283/registration

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY . .
RUN dep ensure
RUN CC=$(which musl-gcc) go build --ldflags '-w -linkmode external -extldflags "-static"' -o main ./cmd/registration

FROM golang:alpine

WORKDIR /app

COPY --from=builder /go/src/github.com/cjhall1283/registration/main .

CMD ["/app/main"]