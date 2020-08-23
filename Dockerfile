FROM golang:1.14

WORKDIR /go/src/chatsapp
COPY . /go/src/chatsapp

RUN go build -o ./bin/chatsapp ./cmd/chatsapp/
# Для возможности запуска скрипта
RUN chmod +x /go/src/chatsapp/scripts/*

EXPOSE 9000/tcp

CMD [ "/go/src/chatsapp/bin/chatsapp" ]



