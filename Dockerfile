FROM golang:1.12.0 AS build
COPY . /go/src/github.com/bborbe/disk-status
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o /main ./src/github.com/bborbe/disk-status/cmd/disk-status-server
CMD ["/bin/bash"]

FROM alpine:3.9 as alpine
RUN apk --no-cache add ca-certificates

FROM scratch
MAINTAINER Benjamin Borbe <bborbe@rocketnews.de>
COPY --from=build /main /main
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/main"]
