FROM registry.access.redhat.com/ubi8/ubi-minimal

RUN mkdir /app
COPY bin/golangy-amd64-linux /app/golangy

ENTRYPOINT ["/app/golangy"]