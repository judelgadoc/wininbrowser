FROM golang:1.19-alpine AS builder

COPY . /interface

WORKDIR /interface

RUN go build -o wininbrowser_int

FROM alpine

WORKDIR /root/

COPY --from=builder /interface .

EXPOSE 55694

CMD ["./wininbrowser_int"]
