FROM alpine:latest
EXPOSE 3306 8080
COPY manifest /manifest
COPY Dockerfile /Dockerfile
COPY bin/database-mesh /database-mesh
WORKDIR /
ENTRYPOINT ["/database-mesh"]
