FROM golang:alpine

WORKDIR /var/www

COPY . .

RUN go install github.com/githubnemo/CompileDaemon@latest

EXPOSE $PORT

ENTRYPOINT ["sh", "/var/www/up.sh"]