FROM golang
COPY ./ /go/src/github.com/joshmenden/planet-express/ship
WORKDIR /go/src/github.com/joshmenden/planet-express/ship

# grpc installation stesp from https://hub.docker.com/r/grpc/go/dockerfile
RUN apt-get update && \
    apt-get -y install git unzip build-essential autoconf libtool
# RUN git clone https://github.com/google/protobuf.git && \
#     cd protobuf && \
#     ./autogen.sh && \
#     ./configure && \
#     make && \
#     make install && \
#     ldconfig && \
#     make clean && \
#     cd .. && \
#     rm -r protobuf

RUN go get google.golang.org/grpc
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go build -o /ship
ENTRYPOINT /ship
EXPOSE 10000