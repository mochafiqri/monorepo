FROM golang AS builder
RUN mkdir /jds-fetch
WORKDIR /jds-fetch
COPY . .
WORKDIR /jds-fetch/cmd
RUN go build -o jds_fetch
RUN chmod +x jds_fetch

FROM debian:buster-slim
RUN apt-get update && apt-get install -y ca-certificates
RUN ln -sf /usr/share/zoneinfo/Asia/Jakarta /etc/localtime
RUN mkdir /jds_fetch
WORKDIR /jds_fetch
COPY --from=builder /jds-fetch/cmd /jds_fetch/cmd
ENTRYPOINT ["/jds_fetch/cmd"]