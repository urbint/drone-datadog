FROM gliderlabs/alpine:3.1
RUN apk-install ca-certificates
ADD drone-datadog /bin/
ENTRYPOINT ["/bin/drone-datadog"]
