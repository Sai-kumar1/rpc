FROM golang:1.18-buster

RUN mkdir /totalitycorp

COPY . /totalitycorp

WORKDIR /totalitycorp

RUN go build -o server server.go

CMD [ "/totalitycorp/server" ]