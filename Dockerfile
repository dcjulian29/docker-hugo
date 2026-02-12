FROM golang:1.26-alpine

ARG VERSION
ARG HUGO_BUILD_TAGS=extended

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GO111MODULE=on

RUN apk update && apk add --no-cache gcc g++ musl-dev git
RUN mkdir -p /go/src/github.com/gohugoio
RUN wget -nv https://github.com/gohugoio/hugo/archive/refs/tags/v${VERSION}.tar.gz
RUN tar -x -C /go/src/github.com/gohugoio -f v${VERSION}.tar.gz
RUN mv /go/src/github.com/gohugoio/hugo-${VERSION} /go/src/github.com/gohugoio/hugo
RUN cd /go/src/github.com/gohugoio/hugo && \
    go install github.com/magefile/mage && \
    mage hugo && \
    mage install

FROM golang:1.26-alpine

COPY --from=0 /go/bin/hugo /usr/bin/hugo

RUN apk update && \
    apk add --no-cache ca-certificates libc6-compat libstdc++ git && \
    hugo version

VOLUME /site
WORKDIR /site
EXPOSE 1313

ENTRYPOINT [ "/usr/bin/hugo" ]
CMD [ "--help" ]
