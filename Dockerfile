FROM node:12-alpine AS front-style-build
WORKDIR /src
ADD ./front-style/package.json .
ADD ./front-style/package-lock.json .
RUN npm install
ADD ./front-style/src ./src
RUN npm run -s build

FROM node:12-alpine AS backoffice-ui-build
WORKDIR /src
ADD ./backoffice-ui/package.json .
ADD ./backoffice-ui/package-lock.json .
RUN npm install
ADD ./backoffice-ui/src ./src
RUN npm run -s build

FROM golang:1.13-alpine AS app-build
ENV GOPROXY direct
RUN apk add --update --no-cache git
WORKDIR /src/app
ADD app/go.mod .
ADD app/go.sum .
RUN go mod download && go get -u github.com/gobuffalo/packr/v2/packr2
ADD ./app/main.go .
ADD ./app/pkg ./pkg
ADD ./app/resources ./resources
WORKDIR /src/front-style/dist
COPY --from=front-style-build /src/dist .
WORKDIR /src/backoffice-ui/dist
COPY --from=backoffice-ui-build /src/dist .
WORKDIR /src/app/pkg/assets
RUN packr2
WORKDIR /src/app
RUN go build -ldflags "-s -w" -o dist/gonference .

FROM alpine:3.11
ENV GIN_MODE release
WORKDIR /app
VOLUME /app/uploads
VOLUME /app/logs
COPY --from=app-build /src/app/dist/gonference .
CMD ./gonference
