FROM golang

WORKDIR /work

RUN apt update -y && apt upgrade -y
