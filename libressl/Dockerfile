FROM ubuntu

WORKDIR /app

RUN apt-get update

RUN apt-get install -y wget automake autoconf git libtool perl make

RUN wget https://ftp.openbsd.org/pub/OpenBSD/LibreSSL/libressl-2.8.3.tar.gz

RUN tar -zxf libressl-2.8.3.tar.gz

WORKDIR libressl-2.8.3

RUN ./configure

RUN make install

RUN ldconfig

WORKDIR /mnt
