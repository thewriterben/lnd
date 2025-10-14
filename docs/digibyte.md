# Lightning Network Support for DigiByte

This document describes the Lightning Network implementation for DigiByte (DGB) in LND. This implementation enables off-chain transactions for DigiByte while leveraging its fast 15-second block times and MultiAlgo consensus mechanism.

## Overview

DigiByte is a UTXO-based blockchain similar to Bitcoin with several key differences:

- **15-second block times** (vs Bitcoin's 10 minutes)
- **MultiAlgo mining** (5 algorithms: SHA256, Scrypt, Groestl, Skein, Qubit)
- **Lower fee structure** (typically 1-10 sats/byte)
- **Different address formats** (starts with 'D' for mainnet)

## Features

- Full Lightning Network protocol support on DigiByte
- Support for DigiByte mainnet, testnet, and regtest
- Optimized parameters for DigiByte's fast block times
- Lower default fees appropriate for DigiByte's fee environment
- Compatible with DigiByte Core (digibyted) via bitcoind RPC interface

## Building LND with DigiByte Support

### Prerequisites

- Go 1.21 or later
- DigiByte Core node (digibyted) synced and running
- Git

### Build Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/thewriterben/lnd.git
   cd lnd
   ```

2. Build LND with DigiByte support:
   ```bash
   ./scripts/build-digibyte.sh
   ```

   This will create two binaries:
   - `lnd-digibyte` - The Lightning Network Daemon
   - `lncli-digibyte` - The command-line interface

## Configuration

### DigiByte Core Setup

First, configure your DigiByte Core node (`digibyte.conf`):

```conf
# Enable RPC server
server=1

# RPC credentials (use strong passwords in production!)
rpcuser=your_rpc_username
rpcpassword=your_rpc_password

# RPC bind address
rpcbind=127.0.0.1
rpcallowip=127.0.0.1

# Enable ZMQ notifications
zmqpubrawblock=tcp://127.0.0.1:28332
zmqpubrawtx=tcp://127.0.0.1:28333

# Transaction index (required for LND)
txindex=1
```

### LND Configuration

Create `~/.lnd/lnd.conf`:

```conf
[Application Options]
alias=My DigiByte Lightning Node
debuglevel=info

[DigiByte]
digibyte.mainnet=true
digibyte.node=bitcoind

# Optimized for DigiByte's 15-second blocks
digibyte.timelockdelta=40
digibyte.defaultchanconfs=6

# Lower fees for DigiByte
digibyte.basefee=100
digibyte.feerate=1

[Bitcoind]
bitcoind.rpchost=localhost:14022
bitcoind.rpcuser=your_rpc_username
bitcoind.rpcpass=your_rpc_password
bitcoind.zmqpubrawblock=tcp://127.0.0.1:28332
bitcoind.zmqpubrawtx=tcp://127.0.0.1:28333
```

See `sample-lnd-digibyte.conf` for a complete configuration example.

## Running LND

### Start LND

For mainnet:
```bash
./lnd-digibyte --digibyte.mainnet --digibyte.node=bitcoind
```

For testnet:
```bash
./lnd-digibyte --digibyte.testnet --digibyte.node=bitcoind
```

For regtest (local development):
```bash
./lnd-digibyte --digibyte.regtest --digibyte.node=bitcoind
```

### Create a Wallet

In a separate terminal:
```bash
./lncli-digibyte create
```

Follow the prompts to create a new wallet. Save your seed phrase securely!

### Basic Operations

Get a new address to receive DigiByte:
```bash
./lncli-digibyte newaddress p2wkh
```

Check your balance:
```bash
./lncli-digibyte walletbalance
```

Connect to a peer:
```bash
./lncli-digibyte connect <pubkey>@<host>:<port>
```

Open a channel:
```bash
./lncli-digibyte openchannel <pubkey> <amount_in_satoshis>
```

List channels:
```bash
./lncli-digibyte listchannels
```

Create an invoice:
```bash
./lncli-digibyte addinvoice --amt <amount_in_satoshis>
```

Pay an invoice:
```bash
./lncli-digibyte payinvoice <payment_request>
```

## Network Parameters

### Mainnet

- **Network Magic**: `0xfac3b6da`
- **Default P2P Port**: `12024`
- **Default RPC Port**: `14022`
- **Address Prefix**: `D` (P2PKH), `S` (P2SH)
- **Bech32 HRP**: `dgb`
- **BIP44 Coin Type**: `20`

### Testnet

- **Network Magic**: `0xfcc1b7dc`
- **Default P2P Port**: `12026`
- **Default RPC Port**: `14023`
- **Address Prefix**: `m` or `n` (P2PKH), `t` (P2SH)
- **Bech32 HRP**: `dgbt`

### Regtest

- **Network Magic**: `0xdab5bffa`
- **Default P2P Port**: `18444`
- **Default RPC Port**: `18443`
- **Address Prefix**: Same as testnet
- **Bech32 HRP**: `dgbrt`

## DigiByte-Specific Optimizations

### Block Time Considerations

DigiByte's 15-second block time offers several advantages for Lightning Network:

1. **Faster Channel Opens**: With 6 confirmations, channels open in ~90 seconds (vs 60 minutes on Bitcoin)
2. **Quicker Force Closes**: Unilateral channel closes resolve faster
3. **Lower Timelock Delta**: Default of 40 blocks = 10 minutes (vs Bitcoin's 80 blocks = ~13 hours)

### Fee Structure

DigiByte typically has lower on-chain fees than Bitcoin:

- **Typical Fee Rate**: 1-10 sats/byte
- **Default Base Fee**: 100 millisatoshis (lower than Bitcoin's 1000)
- **Default Fee Rate**: 1 millionth (same as Bitcoin)

### MultiAlgo Awareness

DigiByte uses 5 different mining algorithms (MultiAlgo). While LND doesn't need to be aware of the specific algorithm used for each block, operators should note:

- Block times can vary slightly between algorithms
- Network difficulty adjusts per-block for each algorithm
- Overall block consistency is maintained through MultiAlgo

## Testing

### Running Tests

Run the test suite:
```bash
./scripts/test-digibyte.sh
```

### Regtest Development

For local development, use regtest mode:

1. Start DigiByte Core in regtest mode
2. Start LND in regtest mode:
   ```bash
   ./lnd-digibyte --digibyte.regtest --digibyte.node=bitcoind
   ```
3. Generate blocks as needed:
   ```bash
   digibyte-cli -regtest generate <num_blocks>
   ```

## Limitations and Future Work

### Current Limitations

1. **Backend Support**: Currently only supports digibyted (via bitcoind interface)
   - No btcd-style backend yet
   - No SPV/Neutrino support yet

2. **Testing**: Limited integration test coverage specific to DigiByte
   - Most Bitcoin tests should work but haven't been extensively validated

3. **Documentation**: Need more real-world examples and deployment guides

### Future Enhancements

1. **SPV Support**: Implement Neutrino/SPV mode for DigiByte
2. **Native Backend**: Consider a dedicated DigiByte backend (not via bitcoind interface)
3. **Performance Metrics**: Collect and publish performance data for DigiByte Lightning
4. **Mobile Support**: Build mobile libraries (lndmobile) for DigiByte
5. **Watchtower Support**: Full testing and validation of watchtower functionality
6. **Submarine Swaps**: Test and validate submarine swaps between DigiByte Lightning and on-chain

## Troubleshooting

### Common Issues

**Problem**: "unable to load RPC credentials for digibyted"
- **Solution**: Ensure `digibyte.conf` has correct RPC credentials and LND config matches

**Problem**: "unable to find coin witness for channel funding"
- **Solution**: Ensure `txindex=1` is set in `digibyte.conf` and DigiByte Core is fully synced

**Problem**: "ZMQ connection failed"
- **Solution**: Check ZMQ endpoints in both `digibyte.conf` and `lnd.conf` match

**Problem**: "channel confirmation taking too long"
- **Solution**: Check if DigiByte Core is synced and blocks are being mined (wait ~90 seconds for 6 confirmations)

### Getting Help

- **GitHub Issues**: https://github.com/thewriterben/lnd/issues
- **DigiByte Community**: https://digibyte.org/
- **LND Documentation**: https://docs.lightning.engineering/

## Security Considerations

1. **Use Strong Passwords**: Always use strong, unique passwords for RPC credentials
2. **Firewall Configuration**: Restrict RPC access to localhost in production
3. **Backup Your Seed**: Store your 24-word seed phrase securely offline
4. **Backup Channel State**: Regularly backup your channel state database
5. **Monitor Your Node**: Keep your DigiByte Core and LND nodes monitored and updated
6. **Start Small**: Test with small amounts before committing significant funds

## Performance Notes

Based on DigiByte's characteristics:

- **Expected Channel Open Time**: ~90 seconds (6 confirmations)
- **Expected Force Close Resolution**: Faster than Bitcoin due to 15s blocks
- **On-chain Fee Savings**: Typically 10-100x lower than Bitcoin
- **Routing Efficiency**: Faster block propagation enables more efficient routing

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This implementation follows the same license as LND (MIT License).

## Acknowledgments

- DigiByte Core developers for the excellent blockchain implementation
- Lightning Labs for the LND implementation
- The DigiByte community for support and testing

## References

- [DigiByte Official Website](https://digibyte.org/)
- [DigiByte GitHub](https://github.com/digibyte/digibyte)
- [Lightning Network Specifications](https://github.com/lightningnetwork/lightning-rfc)
- [LND Documentation](https://docs.lightning.engineering/)
- [SLIP-0044: Registered coin types](https://github.com/satoshilabs/slips/blob/master/slip-0044.md)
