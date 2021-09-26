FROM golang:latest AS build

# ENV GOPATH /go
#
#test

RUN mkdir /go/src/gobackend
WORKDIR /go/src/gobackend
COPY ./ /go/src/gobackend/

RUN go mod init 
RUN go mod tidy 

ENV CGO_ENABLED=0
RUN go build -o gobackend .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/gobackend/gobackend ./
COPY --from=build /go/src/gobackend/.env ./
EXPOSE 8080/tcp
ENTRYPOINT ["./gobackend"]


# COPY main.go .

# ENV CGO_ENABLED=0 //prod ?

# ENV GO111MODULE=on
# COPY . .

# RUN go get -d -v golang.org/x/text/unicode
# RUN go get -d -v golang.org/x/tools/go/packages
# RUN go get -d -v github.com/google/gobackend-cmp/cmp
# RUN go get -d -v ./...
# RUN go install -v ./...
#
# RUN go mod init github/ttaki/gobackend

# RUN go build -o gobackend/cmd

# FROM scratch AS runtime
# ENV GIN_MODE=release
# COPY --from=build /go/src/openapi ./
# EXPOSE 8080/tcp
# ENTRYPOINT ["./gobackend"]
