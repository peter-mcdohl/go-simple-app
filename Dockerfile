FROM alpine:latest
WORKDIR /root/
COPY ./app .
CMD ["./app"]
