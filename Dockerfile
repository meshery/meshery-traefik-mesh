FROM golang:1.15 as bd
WORKDIR /github.com/layer5io/meshery-traefik-mesh
ADD . .
RUN GOPROXY=https://proxy.golang.org GOSUMDB=off go build -ldflags="-w -s" -a -o /meshery-traefik-mesh .
RUN find . -name "*.go" -type f -delete

FROM alpine
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN apk --update add ca-certificates
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

USER appuser
RUN mkdir -p /home/appuser/.kube
RUN mkdir -p /home/appuser/.meshery
WORKDIR /home/appuser
COPY --from=bd /meshery-traefik-mesh /home/appuser
# COPY --from=bd /etc/passwd /etc/passwd
CMD ./meshery-traefik-mesh
