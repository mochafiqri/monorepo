FROM golang AS builder
RUN mkdir /fetch
WORKDIR /fetch
COPY . .
RUN go build -o fetch

FROM debian:buster-slim
RUN apt-get update && apt-get install -y  ca-certificates
RUN ln -sf /usr/share/zoneinfo/Asia/Jakarta /etc/localtime
RUN mkdir /fetch-app
WORKDIR /fetch-app
COPY --from=builder /fetch/fetch /fetch-app/
ENTRYPOINT ./fetch