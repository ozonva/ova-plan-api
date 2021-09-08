FROM golang:1.16-buster AS builder

WORKDIR /app
RUN apt-get update && apt-get install -y protobuf-compiler
COPY . /app/
RUN make prepare

FROM alpine:latest
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY --from=builder /app/bin/ .
CMD ["/app/ova-plan-api"]