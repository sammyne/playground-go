FROM ubuntu:18.04

RUN apt update &&\
    apt install build-essential cmake wget -y &&\
    wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.12.7.linux-amd64.tar.gz &&\
    ln -sf /usr/local/go/bin/* /usr/bin/ &&\
    rm -rf go1.*.tar.gz