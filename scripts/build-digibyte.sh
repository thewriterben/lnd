#!/bin/bash

# Build script for LND with DigiByte support
# This script builds the LND binary with DigiByte integration enabled.

set -e

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Building LND with DigiByte support...${NC}"

# Get the directory of this script
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR="$( cd "${SCRIPT_DIR}/.." && pwd )"

cd "${ROOT_DIR}"

# Build tags for DigiByte support
BUILD_TAGS="autopilotrpc,signrpc,walletrpc,chainrpc,invoicesrpc,routerrpc,watchtowerrpc,monitoring,kvdb_postgres,kvdb_etcd"

# Build LND
echo -e "${YELLOW}Compiling LND...${NC}"
go build -v -tags="${BUILD_TAGS}" -o lnd-digibyte ./cmd/lnd

# Build lncli
echo -e "${YELLOW}Compiling lncli...${NC}"
go build -v -tags="${BUILD_TAGS}" -o lncli-digibyte ./cmd/lncli

echo -e "${GREEN}Build complete!${NC}"
echo ""
echo -e "Binaries created:"
echo -e "  ${GREEN}lnd-digibyte${NC}  - LND daemon with DigiByte support"
echo -e "  ${GREEN}lncli-digibyte${NC} - LND command-line interface"
echo ""
echo -e "To run LND with DigiByte mainnet:"
echo -e "  ${YELLOW}./lnd-digibyte --digibyte.mainnet --digibyte.node=bitcoind${NC}"
echo ""
echo -e "To run LND with DigiByte testnet:"
echo -e "  ${YELLOW}./lnd-digibyte --digibyte.testnet --digibyte.node=bitcoind${NC}"
echo ""
echo -e "To run LND with DigiByte regtest:"
echo -e "  ${YELLOW}./lnd-digibyte --digibyte.regtest --digibyte.node=bitcoind${NC}"
echo ""
