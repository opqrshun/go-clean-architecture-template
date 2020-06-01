FROM golang:1.13 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0

# ENV GO111MODULE=on
# COPY . .

# RUN go get -d -v golang.org/x/text/unicode
# RUN go get -d -v golang.org/x/tools/go/packages
# RUN go get -d -v github.com/google/go-cmp/cmp
# RUN go get -d -v ./...
# RUN go install -v ./...
RUN go mod init github/ttaki/go-clean-architecture

RUN go build -o go-clean-architecture .

# FROM scratch AS runtime
# ENV GIN_MODE=release
# COPY --from=build /go/src/openapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./go-clean-architecture"]
