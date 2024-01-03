FROM ubuntu:22.04

LABEL maintainer="@k33g_org"

# Set environment variables for Go version and user
ARG GO_ARCH=${GO_ARCH}
ARG GO_VERSION=${GO_VERSION}
ARG TINYGO_ARCH=${TINYGO_ARCH}
ARG TINYGO_VERSION=${TINYGO_VERSION}
ARG EXTISM_ARCH=${EXTISM_ARCH}
ARG EXTISM_VERSION=${EXTISM_VERSION}

ARG WASMTIME_VERSION=${WASMTIME_VERSION}
ARG WASMTIME_ARCH=${WASMTIME_ARCH}
ARG WASMEDGE_VERSION=${WASMEDGE_VERSION}

ARG NODE_DISTRO=${NODE_DISTRO}
ARG NODE_VERSION=${NODE_VERSION}

ARG USER_NAME=${USER_NAME}

ARG DEBIAN_FRONTEND=noninteractive


ENV LANG=en_US.UTF-8
ENV LANGUAGE=en_US.UTF-8
ENV LC_COLLATE=C
ENV LC_CTYPE=en_US.UTF-8

# Update the system and install necessary tools
RUN <<EOF
apt-get update 
apt-get install -y curl wget git build-essential xz-utils bat exa software-properties-common htop openssh-server
apt-get install -y nano mc
ln -s /usr/bin/batcat /usr/bin/bat
apt-get -y install hey
EOF

# ------------------------------------
# Install Go
# ------------------------------------
RUN <<EOF
GO_ARCH="${GO_ARCH}"
GO_VERSION="${GO_VERSION}"

wget https://golang.org/dl/go${GO_VERSION}.linux-${GO_ARCH}.tar.gz
tar -xvf go${GO_VERSION}.linux-${GO_ARCH}.tar.gz
mv go /usr/local
rm go${GO_VERSION}.linux-${GO_ARCH}.tar.gz
EOF

# ------------------------------------
# Set Environment Variables for Go
# ------------------------------------
#ENV GOROOT=/usr/local/go
#ENV GOPATH=$HOME/go
#ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/home/${USER_NAME}/go"
ENV GOROOT="/usr/local/go"

RUN <<EOF
go version
go install -v golang.org/x/tools/gopls@latest
go install -v github.com/ramya-rao-a/go-outline@latest
go install -v github.com/stamblerre/gocode@v1.0.0
EOF

# ------------------------------------
# Install TinyGo
# ------------------------------------
RUN <<EOF
TINYGO_ARCH="${TINYGO_ARCH}"
TINYGO_VERSION="${TINYGO_VERSION}"

wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo_${TINYGO_VERSION}_${TINYGO_ARCH}.deb
dpkg -i tinygo_${TINYGO_VERSION}_${TINYGO_ARCH}.deb
rm tinygo_${TINYGO_VERSION}_${TINYGO_ARCH}.deb
EOF


# ------------------------------------
# Install Wasm Runtimes
# ------------------------------------
RUN <<EOF
WASMTIME_VERSION="${WASMTIME_VERSION}"
WASMTIME_ARCH="${WASMTIME_ARCH}"
wget https://github.com/bytecodealliance/wasmtime/releases/download/v${WASMTIME_VERSION}/wasmtime-v${WASMTIME_VERSION}-${WASMTIME_ARCH}.tar.xz
tar -xf wasmtime-v${WASMTIME_VERSION}-${WASMTIME_ARCH}.tar.xz
cd wasmtime-v${WASMTIME_VERSION}-${WASMTIME_ARCH}
mv wasmtime /usr/bin
cd ..
rm -rf wasmtime-v${WASMTIME_VERSION}-${WASMTIME_ARCH}
rm wasmtime-v${WASMTIME_VERSION}-${WASMTIME_ARCH}.tar.xz

WASMEDGE_VERSION="${WASMEDGE_VERSION}"
curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -p /usr/local -v $WASMEDGE_VERSION

#curl https://wazero.io/install.sh | sh
#mv ./bin/wazero /usr/bin/wazero
EOF

# ------------------------------------
# Install Extism CLI
# ------------------------------------
RUN <<EOF
EXTISM_ARCH="${EXTISM_ARCH}"
EXTISM_VERSION="${EXTISM_VERSION}"

wget https://github.com/extism/cli/releases/download/v${EXTISM_VERSION}/extism-v${EXTISM_VERSION}-linux-${EXTISM_ARCH}.tar.gz

tar -xf extism-v${EXTISM_VERSION}-linux-${EXTISM_ARCH}.tar.gz -C /usr/bin
rm extism-v${EXTISM_VERSION}-linux-${EXTISM_ARCH}.tar.gz

extism --version
EOF


# ------------------------------------
# Install NodeJS
# ------------------------------------
RUN <<EOF
NODE_VERSION="${NODE_VERSION}"
NODE_DISTRO="${NODE_DISTRO}"
wget https://nodejs.org/dist/${NODE_VERSION}/node-${NODE_VERSION}-${NODE_DISTRO}.tar.xz
mkdir -p /usr/local/lib/nodejs
tar -xJvf node-$NODE_VERSION-$NODE_DISTRO.tar.xz -C /usr/local/lib/nodejs
rm node-$NODE_VERSION-$NODE_DISTRO.tar.xz
EOF

ENV VERSION="${NODE_VERSION}"
ENV DISTRO="${NODE_DISTRO}"
ENV PATH=/usr/local/lib/nodejs/node-$VERSION-$DISTRO/bin:$PATH


# Create a regular user
RUN useradd -ms /bin/bash ${USER_NAME}

# Set the working directory
WORKDIR /home/${USER_NAME}

# Set the user as the owner of the working directory
RUN chown -R ${USER_NAME}:${USER_NAME} /home/${USER_NAME}

# Switch to the regular user
USER ${USER_NAME}


# ------------------------------------
# Install OhMyBash
# ------------------------------------
RUN <<EOF
bash -c "$(curl -fsSL https://raw.githubusercontent.com/ohmybash/oh-my-bash/master/tools/install.sh)"
EOF

# ------------------------------------
# Install Rust + Wasm Toolchain
# ------------------------------------
RUN <<EOF
apt install -y pkg-config libssl-dev
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
export RUSTUP_HOME=~/.rustup
export CARGO_HOME=~/.cargo
export PATH=$PATH:$CARGO_HOME/bin
curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh 
rustup target add wasm32-wasi
EOF

# Set the entrypoint to an interactive shell
CMD ["/bin/bash"]
