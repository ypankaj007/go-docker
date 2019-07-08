FROM golang
COPY . /go/src/go-docker
WORKDIR /go/src/go-docker
RUN go get .; go build -o go-docker
ENTRYPOINT ["./go-docker"]
