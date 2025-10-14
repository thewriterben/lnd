#!/bin/bash

# Test script for LND with DigiByte support
# This script runs tests related to DigiByte functionality

set -e

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Running LND tests with DigiByte support...${NC}"

# Get the directory of this script
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR="$( cd "${SCRIPT_DIR}/.." && pwd )"

cd "${ROOT_DIR}"

# Build tags for DigiByte support
BUILD_TAGS="autopilotrpc,signrpc,walletrpc,chainrpc,invoicesrpc,routerrpc,watchtowerrpc,monitoring,kvdb_postgres,kvdb_etcd"

# Run unit tests for chain parameters
echo -e "${YELLOW}Testing chain parameters...${NC}"
go test -v -tags="${BUILD_TAGS}" ./chainparams/...

# Run unit tests for chainreg
echo -e "${YELLOW}Testing chain registry...${NC}"
go test -v -tags="${BUILD_TAGS}" ./chainreg/...

# Run unit tests for keychain
echo -e "${YELLOW}Testing keychain...${NC}"
go test -v -tags="${BUILD_TAGS}" ./keychain/...

# Run config tests
echo -e "${YELLOW}Testing configuration...${NC}"
go test -v -tags="${BUILD_TAGS}" -run TestConfig ./

echo -e "${GREEN}All tests passed!${NC}"
