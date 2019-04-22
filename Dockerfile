FROM alpine:latest
ADD user-service /app/user-service
ENTRYPOINT [ "/app/user-service" ]
