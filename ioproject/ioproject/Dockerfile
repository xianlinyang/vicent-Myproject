FROM golang:1.17.4-alpine3.15 as builder

WORKDIR $GOPAT/src/ioproject
ADD . ./
ENV GOMODULE=on
RUN go build -o ioproject .



EXPOSE 8084:8084
CMD ["./ioproject"]







