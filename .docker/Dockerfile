FROM gitpod/workspace-base

LABEL maintainer="@k33g_org"

USER root
# ------------------------------------
# Set environment variables
# ------------------------------------
ARG GO_ARCH=${GO_ARCH}
ARG GO_VERSION=${GO_VERSION}
ARG TINYGO_ARCH=${TINYGO_ARCH}
ARG TINYGO_VERSION=${TINYGO_VERSION}
ARG EXTISM_ARCH=${EXTISM_ARCH}
ARG EXTISM_VERSION=${EXTISM_VERSION}
ARG EXTISM_JS_VERSION=${EXTISM_JS_VERSION}
ARG EXTISM_JS_ARCH=${EXTISM_JS_ARCH}
ARG EXTISM_JS_OS=${EXTISM_JS_OS}
ARG NODE_MAJOR=${NODE_MAJOR}
ARG BINARYEN_VERSION=${BINARYEN_VERSION}
ARG BINARYEN_ARCH=${BINARYEN_ARCH}
ARG BINARYEN_OS=${BINARYEN_OS}

ARG HELM_VERSION=${HELM_VERSION}
ARG HELM_ARCH=${HELM_ARCH}

ARG K9S_VERSION=${K9S_VERSION}
ARG K9S_ARCH=${K9S_ARCH}

ARG DEBIAN_FRONTEND=noninteractive

ENV LANG=en_US.UTF-8
ENV LANGUAGE=en_US.UTF-8
ENV LC_COLLATE=C
ENV LC_CTYPE=en_US.UTF-8

# Update the system and install necessary tools
RUN <<EOF
apt-get update 
apt-get install -y curl wget git build-essential xz-utils bat exa software-properties-common htop openssh-server sshpass gettext
apt install -y pkg-config libssl-dev
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
apt-get update && apt-get install -y ca-certificates curl gnupg
curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg
echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" | tee /etc/apt/sources.list.d/nodesource.list
apt-get update && apt-get install nodejs -y
EOF

# -----------------------
# Install Extism JS PDK
# -----------------------

RUN <<EOF
export EXTISM_JS_VERSION="${EXTISM_JS_VERSION}"
export EXTISM_JS_ARCH="${EXTISM_JS_ARCH}"
export EXTISM_JS_OS="${EXTISM_JS_OS}"
curl -L -O "https://github.com/extism/js-pdk/releases/download/v${EXTISM_JS_VERSION}/extism-js-${EXTISM_JS_ARCH}-${EXTISM_JS_OS}-v${EXTISM_JS_VERSION}.gz"
gunzip extism-js*.gz
chmod +x extism-js-*
mv extism-js-* /usr/local/bin/extism-js
EOF

# ------------------------------------
# Install Wasmtime, Wazero, Wasmer CLI
# ------------------------------------
RUN <<EOF
curl https://wasmtime.dev/install.sh -sSf | bash

curl https://wazero.io/install.sh | sh
mv ./bin/wazero /usr/bin/wazero

curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash
EOF

# ------------------------------------
# Install Rust + Wasm Toolchain
# ------------------------------------
#RUN <<EOF
#apt install -y pkg-config libssl-dev
#curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
#export RUSTUP_HOME=~/.rustup
#export CARGO_HOME=~/.cargo
#export PATH=$PATH:$CARGO_HOME/bin
#curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh 
#rustup target add wasm32-wasi
#EOF

# ------------------------------------
# Install Helm
# ------------------------------------
RUN <<EOF
wget https://get.helm.sh/helm-v${HELM_VERSION}-linux-${HELM_ARCH}.tar.gz
tar xvf helm-v${HELM_VERSION}-linux-${HELM_ARCH}.tar.gz
mv linux-${HELM_ARCH}/helm /usr/local/bin
EOF

# ------------------------------------
# Install K8S tools: kubectl & k9s
# ------------------------------------
RUN <<EOF
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

chmod +x kubectl
mv ./kubectl /usr/local/bin/kubectl

wget https://github.com/derailed/k9s/releases/download/v${K9S_VERSION}/k9s_${K9S_ARCH}.tar.gz
tar xvf k9s_${K9S_ARCH}.tar.gz -C /usr/local/bin
EOF


USER gitpod

# ------------------------------------
# Install Binaryen
# ------------------------------------
RUN <<EOF

wget https://github.com/WebAssembly/binaryen/releases/download/${BINARYEN_VERSION}/binaryen-${BINARYEN_VERSION}-${BINARYEN_ARCH}-${BINARYEN_OS}.tar.gz
tar -xf binaryen-${BINARYEN_VERSION}-${BINARYEN_ARCH}-${BINARYEN_OS}.tar.gz

sudo cp binaryen-${BINARYEN_VERSION}/bin/* /usr/bin
rm -rf binaryen-${BINARYEN_VERSION}
rm binaryen-${BINARYEN_VERSION}-${BINARYEN_ARCH}-${BINARYEN_OS}.tar.gz

wasm2js --version
EOF

# ------------------------------------
# Set Environment Variables for Go
# ------------------------------------
ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

RUN <<EOF
go version
go install -v golang.org/x/tools/gopls@latest
go install -v github.com/ramya-rao-a/go-outline@latest
go install -v github.com/stamblerre/gocode@v1.0.0
EOF

# ------------------------------------
# Install Rust + Wasm Toolchain
# ------------------------------------
RUN <<EOF
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
export RUSTUP_HOME=~/.rustup
export CARGO_HOME=~/.cargo
export PATH=$PATH:$CARGO_HOME/bin
curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh 
rustup target add wasm32-wasi
EOF

# Command to run when starting the container
CMD ["/bin/bash"]
