FROM node:12.16-slim AS JS_BUILD
COPY webapp /webapp
WORKDIR /webapp
RUN yarn && yarn build

FROM golang:1.14-alpine AS GO_BUILD
RUN apk update && apk --no-cache add build-base
COPY server /server
WORKDIR /server
RUN make clean build

FROM alpine:3.11
COPY --from=JS_BUILD /webapp/dist/* ./webapp/
COPY --from=GO_BUILD /server/bin/server ./
CMD ./server

LABEL maintainer="Stas Lebedev <kugui@yandex.ru>"
