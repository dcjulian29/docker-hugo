FROM golang:latest

ARG VERSION
ARG HUGO_BUILD_TAGS=extended

RUN cd /usr/local/src \
  && wget -nv https://github.com/gohugoio/hugo/archive/refs/tags/v${VERSION}.tar.gz \
  && tar -xvf v${VERSION}.tar.gz \
  && cd hugo-${VERSION} \
  && go build -o /usr/local/src/hugo main.go

#---------------------------------------------

FROM debian:11-slim

COPY --from=0 /usr/local/src/hugo /usr/bin/hugo

VOLUME /site
WORKDIR /site
EXPOSE 1313

ENTRYPOINT [ "hugo" ]
CMD [ "--help" ]
