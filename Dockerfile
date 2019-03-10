FROM golang:1.11

WORKDIR /go/src/detoplan-go
COPY . .

RUN apt-get update -qq && \
    apt-get install -y mysql-client vim && \
    go get -u github.com/golang/dep/cmd/dep && \
    go get -u github.com/oxequa/realize && \
    go get -u github.com/swaggo/swag/cmd/swag && \
    go get -u github.com/davecgh/go-spew/spew && \
    dep ensure -v -vendor-only

EXPOSE 1323

CMD ["realize", "start"]
