# Grab the latest alpine image
FROM golang:1.19-alpine as build

RUN apk --update add libusb libusb-dev pkgconfig gcc linux-headers musl-dev
COPY . /app
WORKDIR /app

RUN go mod download
RUN mkdir /app/bin
RUN GOOS=linux go build -o /app/bin/goodTimer .


FROM golang:1.19-alpine

RUN apk --update add libusb
COPY --from=build /app /app

CMD ["/app/bin/goodTimer", "-f", "/app/server/etc/goodtimer-api.yaml"]