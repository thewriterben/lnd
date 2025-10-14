# DigiByte Lightning Network Support

This fork adds Lightning Network support for DigiByte (DGB) to LND.

## Quick Links

📚 **[Full DigiByte Documentation](digibyte.md)** - Complete guide for using LND with DigiByte

## Quick Start

### Build
```bash
./scripts/build-digibyte.sh
```

### Run
```bash
# Mainnet
./lnd-digibyte --digibyte.mainnet --digibyte.node=bitcoind

# Testnet
./lnd-digibyte --digibyte.testnet --digibyte.node=bitcoind
```

### Create Wallet
```bash
./lncli-digibyte create
```

## Key Features

✅ Full Lightning Network protocol support on DigiByte  
✅ Optimized for DigiByte's 15-second block times  
✅ Lower fees appropriate for DigiByte's fee environment  
✅ Support for mainnet, testnet, and regtest  
✅ Compatible with DigiByte Core (digibyted)  

## What's Different from Bitcoin?

| Feature | Bitcoin | DigiByte |
|---------|---------|----------|
| Block Time | 10 minutes | 15 seconds |
| Channel Confirmations | ~60 minutes | ~90 seconds |
| Timelock Delta | 80 blocks (~13 hours) | 40 blocks (~10 minutes) |
| Typical Fees | 10-100 sats/byte | 1-10 sats/byte |
| Base Fee | 1000 msat | 100 msat |

## Requirements

- Go 1.21 or later
- DigiByte Core node (digibyted) with:
  - RPC enabled
  - ZMQ notifications enabled
  - Transaction index enabled (`txindex=1`)

## Sample Configuration

See [sample-lnd-digibyte.conf](../sample-lnd-digibyte.conf) for a complete configuration example.

## Testing

```bash
# Run DigiByte-specific tests
./scripts/test-digibyte.sh

# Run all tests
go test ./...
```

## Documentation

- [Complete DigiByte Guide](digibyte.md) - Comprehensive documentation
- [Sample Configuration](../sample-lnd-digibyte.conf) - Configuration template
- [Build Script](../scripts/build-digibyte.sh) - Build automation
- [Test Script](../scripts/test-digibyte.sh) - Test automation

## Support

- **Issues**: [GitHub Issues](https://github.com/thewriterben/lnd/issues)
- **DigiByte Community**: https://digibyte.org/
- **LND Documentation**: https://docs.lightning.engineering/

## License

MIT License (same as LND)

## Acknowledgments

- Lightning Labs for the LND implementation
- DigiByte Core developers
- The DigiByte community
