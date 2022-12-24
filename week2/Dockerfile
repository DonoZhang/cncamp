FROM scratch

ARG version

COPY bin/amd64/httpserver /httpserver
ENV VERSION=$version
EXPOSE 8080

ENTRYPOINT ["/httpserver"]