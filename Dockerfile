FROM golang:1.12-alpine AS build
RUN apk add --update nodejs npm curl git
WORKDIR /usr
RUN curl -sL https://taskfile.dev/install.sh | sh
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go get
COPY cmd ./cmd
COPY src ./src
COPY Taskfile.yml .
RUN task install-dependencies && \
    task backoffice-ui-build && \
    task build

FROM alpine:3.9
WORKDIR /app
COPY --from=build /app/out/gonference .
CMD ./gonference
