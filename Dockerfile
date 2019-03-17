FROM golang:1.11

WORKDIR /go/src/hashagime
COPY . .

RUN apt-get update -qq && \
    apt-get install -y mysql-client vim p7zip-full && \
    go get -u github.com/golang/dep/cmd/dep && \
    go get -u github.com/oxequa/realize && \
    go get -u github.com/swaggo/swag/cmd/swag && \
    go get -u github.com/davecgh/go-spew/spew && \
    dep ensure -v -vendor-only && \
    wget https://johnvansickle.com/ffmpeg/builds/ffmpeg-git-amd64-static.tar.xz && \
    7z x ffmpeg-git-amd64-static.tar.xz && \
    7z x ffmpeg-git-amd64-static.tar && \
    cp ffmpeg-git-20190305-amd64-static/ffmpeg /usr/bin/  && \
    chmod 755 /usr/bin/ffmpeg


EXPOSE 1323

CMD ["realize", "start"]
