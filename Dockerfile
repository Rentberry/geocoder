# Service
FROM alpine:latest AS runtime

RUN apk --no-cache add ca-certificates

WORKDIR /root
COPY /geocoder .
USER root
CMD ["./geocoder"]

EXPOSE 8080 9092
