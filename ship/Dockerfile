FROM golang
COPY ./ /go/src/github.com/joshmenden/planet-express/ship
WORKDIR /go/src/github.com/joshmenden/planet-express/ship
RUN go get google.golang.org/grpc
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go build -o /ship
ENTRYPOINT /ship
EXPOSE 10000