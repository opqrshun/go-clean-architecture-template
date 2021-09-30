FROM golang:latest AS build


WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./model ./model
COPY ./pkg ./pkg


ENV CGO_ENABLED=0
RUN go build -o gobackend ./cmd

# runtime stage
FROM scratch AS runtime

WORKDIR /app

ENV GIN_MODE=release
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/gobackend /app/gobackend

EXPOSE 8080/tcp
ENTRYPOINT ["/app/gobackend"]


