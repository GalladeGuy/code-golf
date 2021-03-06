FROM debian:stretch-slim

ENV CGO_ENABLED=0 GOROOT_BOOTSTRAP=/usr/local/go

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates curl git make nasm

# https://golang.org/dl/
RUN curl -SSL https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz \
  | tar -xzC /usr/local

RUN git clone https://go.googlesource.com/go \
 && cd go                                    \
 && git checkout ee55f08                     \
 && cd src                                   \
 && ./make.bash                              \
 && chmod +rx /root

ENV GOCACHE=/code-golf/.cache GOPATH=/root/go PATH=/go/bin:$PATH

COPY go.mod go.sum /

RUN go mod download && chmod -R o+w /root/go

CMD go build -ldflags -s                           \
 && nasm -f bin -o run-container run-container.asm \
 && chmod +x run-container
