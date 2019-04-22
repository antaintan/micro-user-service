FROM alpine:latest
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN apk add tzdata
ADD user-service /app/user-service
ENTRYPOINT [ "/app/user-service" ]
