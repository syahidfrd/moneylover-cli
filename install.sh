#!/bin/sh
set -e

REPO="syahidfrd/moneylover-cli"
BINARY="moneylover"
INSTALL_DIR="/usr/local/bin"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH" >&2; exit 1 ;;
esac

case "$OS" in
  linux) OS="linux" ;;
  darwin) OS="darwin" ;;
  *) echo "Unsupported OS: $OS" >&2; exit 1 ;;
esac

FILENAME="moneylover-cli_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/latest/download/${FILENAME}"

TMPDIR=$(mktemp -d)
trap 'rm -rf "$TMPDIR"' EXIT

echo "Downloading ${BINARY} for ${OS}/${ARCH}..."
curl -sSL "$URL" -o "$TMPDIR/$FILENAME"

echo "Extracting..."
tar -xzf "$TMPDIR/$FILENAME" -C "$TMPDIR"

echo "Installing to ${INSTALL_DIR}..."
if [ -w "$INSTALL_DIR" ]; then
  mv "$TMPDIR/$BINARY" "$INSTALL_DIR/$BINARY"
else
  sudo mv "$TMPDIR/$BINARY" "$INSTALL_DIR/$BINARY"
fi

chmod +x "$INSTALL_DIR/$BINARY"

echo "Done! Run 'moneylover --help' to get started."
