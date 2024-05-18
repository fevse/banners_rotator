# TODO
FROM golang:1.22 as build

ENV BIN_FILE /opt/rotator/rotator-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
COPY . /go/src/
RUN CGO_ENABLED=0 go build \
    -ldflags "$LDFLAGS" \
    -o ${BIN_FILE} cmd/rotator/*

FROM alpine:3.19

LABEL ORGANIZATION="fevse"
LABEL SERVICE="rotator"
LABEL MAINTEINERS="fevse"

ENV BIN_FILE "/opt/rotator/rotator-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/rotator/config.toml
COPY ./configs/config.toml ${CONFIG_FILE}

ENV CONFIG_MIGRATION /etc/migrations/
COPY ./migrations ${CONFIG_MIGRATION}

CMD ${BIN_FILE} -config ${CONFIG_FILE}
