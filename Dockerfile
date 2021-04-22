# Asset build
FROM node:12 as assetBuild

COPY ./pkg/http/web /usr/src/app/pkg/http/web/

WORKDIR /usr/src/app/pkg/http/web
RUN npm ci
RUN ./node_modules/.bin/vue-cli-service build

# App build
FROM golang:1.16-alpine as appBuild

COPY ./ /usr/src/app

WORKDIR /usr/src/app
RUN go build -o sep ./cmd/stream_event_proxy/main.go

# Image build
FROM alpine

RUN addgroup -g 2001 sep \
    && adduser -DHu 1001 -G sep sep

COPY --from=appBuild /usr/src/app/sep /app/
COPY --from=appBuild /usr/src/app/web /app/web/
COPY --from=assetBuild /usr/src/app/web/assets /app/web/assets/

WORKDIR /app
USER sep

CMD ["./sep"]