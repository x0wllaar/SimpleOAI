FROM ubuntu:22.04

#GENERIC DOOM EMACS BASED ENVRIONMENT

RUN apt-get update -y && \
    apt-get install -y gpg ca-certificates && \
    echo -n 'deb https://ppa.launchpadcontent.net/kelleyk/emacs/ubuntu jammy main' > /etc/apt/sources.list.d/ondrej-ubuntu-php-groovy.list && \
    apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 873503A090750CDAEB0754D93FF0E01EEAAFC9CD && \
    apt-get update -y && \
    apt-get install -y emacs28 git ripgrep fonts-symbola fonts-cascadia-code wget xz-utils libvterm-dev cmake build-essential && \
    cd /usr/share/fonts && \
    wget https://github.com/ryanoasis/nerd-fonts/releases/download/v3.0.2/FiraCode.tar.xz && \
    wget https://github.com/ryanoasis/nerd-fonts/releases/download/v3.0.2/CascadiaCode.tar.xz && \
    tar xvf FiraCode.tar.xz && \
    tar xvf CascadiaCode.tar.xz && \
    rm FiraCode.tar.xz CascadiaCode.tar.xz && \
    fc-cache -f -v

COPY ./emacs_runner.sh /bin/emacs_runner
COPY ./doom_golang_config /root/default_doom_conf

#GOLANG SPECIFIC STUFF

RUN mkdir /tmp/golang && cd /tmp/golang && \
    wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz
ENV PATH="$PATH:/usr/local/go/bin:/root/go/bin"

RUN go install golang.org/x/tools/gopls@latest && \
    go install mvdan.cc/gofumpt@latest

CMD ["/bin/emacs_runner"]