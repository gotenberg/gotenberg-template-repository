ARG GOLANG_VERSION
ARG GOTENBERG_VERSION

FROM golang:$GOLANG_VERSION AS builder

ENV CGO_ENABLED=0

# Define the working directory outside of $GOPATH (we're using go modules).
WORKDIR /home

# Install module dependencies.
COPY go.mod go.sum ./

RUN go mod download &&\
    go mod verify

# Copy the source code.
COPY cmd ./cmd
COPY pkg ./pkg

# Build the binary.
ARG GOTENBERG_VERSION
ARG APP_NAME
ARG APP_VERSION

RUN go build -o gotenberg -ldflags "-X 'github.com/gotenberg/gotenberg/v8/cmd.Version=$GOTENBERG_VERSION-$APP_NAME-$APP_VERSION'" cmd/app/main.go

FROM gotenberg/gotenberg:$GOTENBERG_VERSION

ARG GOTENBERG_VERSION
ARG APP_AUTHOR
ARG APP_REPOSITORY
ARG APP_NAME
ARG APP_VERSION

LABEL org.opencontainers.image.title="Gotenberg $APP_NAME" \
      org.opencontainers.image.description="A containerized API for seamless PDF conversion." \
      org.opencontainers.image.version="$GOTENBERG_VERSION-$APP_NAME-$APP_VERSION" \
      org.opencontainers.image.authors="$APP_AUTHOR" \
      org.opencontainers.image.source="$APP_REPOSITORY"

USER root

# TODO: add your instructions.

# Copy the Gotenberg binary from the builder stage.
COPY --from=builder /home/gotenberg /usr/bin/

USER gotenberg