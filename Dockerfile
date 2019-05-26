FROM golang:1.12-alpine AS build
RUN apk add --update nodejs npm curl git
RUN go get -u github.com/sirikon/tsk/cmd/tsk
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go get
COPY cmd ./cmd
COPY src ./src
COPY scripts ./scripts
COPY Tskfile.yml .
RUN tsk install
RUN tsk backoffice build
RUN tsk front-style build
RUN tsk build

FROM alpine:3.9
WORKDIR /app
COPY --from=build /app/out/gonference .
CMD ./gonference
