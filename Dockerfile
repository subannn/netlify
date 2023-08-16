FROM golang:alpine

WORKDIR /app

COPY . .

ENV ENDPOINT =          192.168.0.106:9000      
ENV ACCESS_KEY_ID =     TDIF7k0TmCXXClbV1fil
ENV SECRET_ACCESS_KEY = XSFbrmbtyKKGOI5QzlITE9tGpQ9Peagu4ON5FkA4


RUN go build main.go

ENTRYPOINT /app/main

