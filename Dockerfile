FROM alpine
RUN mkdir -p /app/configs/
#RUN  CGO_ENABLED=0  GOOS=linux  GOARCH=amd64  go build -o blog-service main.go
COPY blog-service /app
WORKDIR /app
COPY configs/app/config_release.yaml /app/configs/config_release.yaml
#COPY ../templates /app/templates
RUN chmod 777 /app/blog-service
EXPOSE 8080
ENV go_env=release
ENTRYPOINT ["/app/blog-service"]
