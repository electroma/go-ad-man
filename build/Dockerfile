FROM golang

WORKDIR /tmp
RUN wget http://ftp.us.debian.org/debian/pool/main/s/swig/swig3.0_3.0.10-1.1_amd64.deb && \
    dpkg -i swig3.0_3.0.10-1.1_amd64.deb && \
    ln -s /usr/bin/swig3.0 /usr/bin/swig && \
    apt-get update && \
    apt-get install -y libldap2-dev libkrb5-dev libsasl2-dev