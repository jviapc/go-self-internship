FROM golang:1.24.0-alpine

ARG HOST_UID=1000

RUN set -xeu \
    && addgroup -g ${HOST_UID} -S appgroup && adduser -u ${HOST_UID} -S appuser -G appgroup \
    && install -d -o appuser -g appgroup /app

WORKDIR /app
