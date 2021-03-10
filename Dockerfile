#Compile stage
FROM golang:latest AS build-env

ADD .dockerdev
WORKDIR dockerdev

RUN go build -o /server

# Final stage
FROM alpine:buster

EXPOSE 8000

WORKDIR /
COPY --from build-end /server /

CMD ["/server"]