FROM golang:1.24.3-alpine AS builder

RUN apk --update --no-cache add \
    binutils \
    make \
    && rm -rf /root/.cache
WORKDIR /go/src/github.com/jhoblitt/expfake
COPY . .
RUN make build

FROM alpine:3
WORKDIR /root/
COPY --from=builder /go/src/github.com/jhoblitt/expfake/expfake /bin/expfake
ENTRYPOINT ["/bin/expfake"]
