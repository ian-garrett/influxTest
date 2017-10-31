FROM golang:1.9

RUN apt-get -y update --fix-missing
RUN apt-get -y upgrade
RUN go get github.com/bugsnag/bugsnag-go
RUN go get github.com/influxdata/influxdb/client/v2