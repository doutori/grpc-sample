# Build
FROM golang:1.11.0

RUN mkdir -p /go/src \
    && mkdir -p /go/bin \
    && mkdir -p /go/pkg

ARG src_path

RUN mkdir -p $GOPATH/$src_path
ADD . $GOPATH/$src_path
WORKDIR $GOPATH/$src_path
# RUN apt-get update \
#     && apt-get -y install curl jq \
#     && go get -u github.com/oxequa/realize \
#     && go get -u github.com/golang/dep/cmd/dep \
#     && dep ensure -v -vendor-only=true \

EXPOSE 8080
CMD ["/bin/bash"]
# CMD ["/go/bin/api"]
