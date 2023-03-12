FROM golang:1.19-alpine AS builder

COPY . /clockms

WORKDIR /clockms

RUN go build -o wininbrowser_clock_ms

FROM alpine

RUN apk add --no-cache tzdata

WORKDIR /root/

COPY --from=builder /clockms .

EXPOSE 9090

CMD ["./wininbrowser_clock_ms"]
