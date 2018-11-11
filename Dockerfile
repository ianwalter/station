FROM golang:1.11.2 as build
WORKDIR /go/src/github.com/appjumpstart/station/
# RUN go get -d -v golang.org/x/net/html
COPY app.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
COPY --from=build /go/src/github.com/appjumpstart/station/station /usr/bin/station
