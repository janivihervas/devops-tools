FROM alpine:3.6

RUN apk --no-cache add ca-certificates

WORKDIR app

COPY server .

ENTRYPOINT /app/server

EXPOSE 3000