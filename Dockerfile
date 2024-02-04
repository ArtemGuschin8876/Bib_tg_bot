
FROM golang:1.21.1 AS base

WORKDIR /app

COPY . /app

RUN apt update && apt install  nano curl  net-tools -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN go build -o bib_bot

FROM bitnami/minideb:latest

WORKDIR /app

RUN apt update && apt install nano curl -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY --from=base /app/bib_bot /app

COPY --from=base /app/.env  /app/.env

COPY --from=base /app/translation-api-412323-7e71139aef69.json /app/translation-api-412323-7e71139aef69.json
