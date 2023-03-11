FROM golang:1.19-alpine AS builder

COPY . /clockms

WORKDIR /clockms

RUN go build -o wininbrowser_clock_ms

FROM alpine

WORKDIR /root/

COPY --from=builder /clockms .

EXPOSE 9090

ENV DBHOST=
ENV DBUSER=
ENV DBPASS=

CMD ["./wininbrowser_clock_ms"]
