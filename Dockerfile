
FROM golang:1.21.1 AS base

WORKDIR /app

COPY . /app

RUN apt update && apt install  nano curl  net-tools -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN go build

FROM bitnami/minideb:latest

WORKDIR /app

RUN apt update && apt install nano curl -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY --from=base /app/Bib_bot /app

COPY --from=base /app/.env  /app/.env
