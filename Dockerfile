# Compile stage
FROM debian:buster

# libxslt1 missing in this release
RUN apt update && \
    apt install -y libxml2 libxslt1-dev liblzma5 zlib1g

EXPOSE 8090

WORKDIR /
COPY  /test-go-xslt /server

CMD ["/server"]

