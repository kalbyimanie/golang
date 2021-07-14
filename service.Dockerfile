FROM alpine:edge
COPY ./src/main/main /opt/bin/main
CMD [ "/bin/sh", "-c", "/opt/bin/main" ]