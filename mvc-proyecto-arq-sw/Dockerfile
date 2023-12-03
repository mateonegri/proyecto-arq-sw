FROM golang:latest

ADD . /mvc-proyecto-arq-sw

WORKDIR /mvc-proyecto-arq-sw
RUN go mod init mvc-proyecto-arq-sw
RUN go mod tidy
RUN go build -o mvc-proyecto-arq-sw .
RUN chmod +x /mvc-proyecto-arq-sw

EXPOSE 8090

ENTRYPOINT ["./mvc-proyecto-arq-sw"]