FROM golang:1.14

# don't place it into $GOPATH/bin because Drone mounts $GOPATH as volume
COPY goreleaser-app /usr/bin/
#CMD ["./goreleaser-app version"]