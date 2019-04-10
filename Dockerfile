FROM golang:1.11.0 as build
WORKDIR /go/src/github.com/terrylowery/slug-sync
RUN go get -d -v github.com/rackspace/gophercloud
COPY *.go /go/src/github.com/terrylowery/slug-sync/
RUN CGO_ENABLED=0 GOOS=linux go build -a -v -o ./bin/slug-sync .





FROM ubuntu:trusty

ENV PATH=$PATH:/app

RUN apt-get update && \
    apt-get upgrade && apt-get install -y ca-certificates



RUN mkdir -p /var/cache/chef

WORKDIR /app
COPY --from=build /go/src/github.com/terrylowery/slug-sync/bin/slug-sync .
ADD slug-config.json /etc/slug-sync/

CMD ["bash"]




