FROM golang:1.11.2 as build
WORKDIR /go/src/github.com/appjumpstart/station/
# RUN go get -d -v golang.org/x/net/html
COPY app.go .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=build /go/src/github.com/appjumpstart/station/station .
