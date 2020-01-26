FROM node:12-alpine AS front-style-build
RUN apk add --update --no-cache bash
WORKDIR /src
ADD ./src/front-style/tasks.sh .
ADD ./src/front-style/package.json .
ADD ./src/front-style/package-lock.json .
RUN ./tasks.sh install
ADD ./src/front-style/src ./src
RUN ./tasks.sh build

FROM node:12-alpine AS backoffice-ui-build
RUN apk add --update --no-cache bash
WORKDIR /src
ADD ./src/backoffice-ui/tasks.sh .
ADD ./src/backoffice-ui/package.json .
ADD ./src/backoffice-ui/package-lock.json .
RUN ./tasks.sh install
ADD ./src/backoffice-ui/src ./src
RUN ./tasks.sh build

FROM golang:1.13-alpine AS app-build
ENV GOPROXY direct
RUN apk add --update --no-cache bash git
WORKDIR /src/app
ADD ./src/app/tasks.sh .
ADD ./src/app/go.mod .
ADD ./src/app/go.sum .
RUN ./tasks.sh install
ADD ./src/app/main.go .
ADD ./src/app/pkg ./pkg
ADD ./src/app/resources ./resources
WORKDIR /src/front-style/dist
COPY --from=front-style-build /src/dist .
WORKDIR /src/backoffice-ui/dist
COPY --from=backoffice-ui-build /src/dist .
WORKDIR /src/app
RUN ./tasks.sh packAssets && ./tasks.sh build

FROM alpine:3.11
ENV GIN_MODE release
WORKDIR /app
VOLUME /app/uploads
VOLUME /app/logs
COPY --from=app-build /src/app/dist/gonference .
CMD ./gonference
