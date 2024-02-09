#!/bin/bash

REPO="https://github.com/vistormu/proxima/releases/latest/download"
PROGRAM="proxima"
EXT=""

# Detect the platform
OS="$(uname -s)"
ARCH="$(uname -m)"

if [ "$OS" = "Linux" ]; then
    EXT="-linux-$ARCH"
elif [ "$OS" = "Darwin" ]; then
    EXT="-darwin-$ARCH"
fi

URL="${REPO}/${PROGRAM}${EXT}"

# Download and install
wget -O /tmp/${PROGRAM} ${URL}
chmod +x /tmp/${PROGRAM}
mv /tmp/${PROGRAM} /usr/local/bin/${PROGRAM}

echo "${PROGRAM} installed successfully!"
