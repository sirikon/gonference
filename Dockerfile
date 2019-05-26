FROM golang:1.12-alpine AS build
RUN apk add --update nodejs npm curl git
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go get
COPY cmd ./cmd
COPY src ./src
COPY scripts ./scripts
RUN sh ./scripts/install.sh
ENV BACKOFFICE_UI_PATH=src/backoffice-ui
RUN sh ./scripts/backoffice/build.sh
RUN sh ./scripts/build.sh

FROM alpine:3.9
WORKDIR /app
COPY --from=build /app/out/gonference .
CMD ./gonference
