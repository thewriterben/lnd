// Package chainparams defines chain configuration parameters for DigiByte.
//
// DigiByte is a UTXO blockchain similar to Bitcoin with the following
// key differences:
//   - 15-second block times (vs Bitcoin's 10 minutes)
//   - MultiAlgo mining (5 algorithms: SHA256, Scrypt, Groestl, Skein, Qubit)
//   - Lower fee structure (1-10 sats/byte typical)
//   - Different address prefixes (starts with 'D' for mainnet)
package chainparams

import (
	bitcoinCfg "github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// DigiByteMainNetParams defines the network parameters for the main DigiByte
// network.
var DigiByteMainNetParams = bitcoinCfg.Params{
	Name:        "digibyte-mainnet",
	Net:         wire.BitcoinNet(0xfac3b6da), // DigiByte mainnet magic
	DefaultPort: "12024",

	// Human-readable part for Bech32 encoded segwit addresses, as defined
	// in BIP 173.
	Bech32HRPSegwit: "dgb",

	// Address encoding magics
	PubKeyHashAddrID:        0x1e, // Starts with 'D'
	ScriptHashAddrID:        0x3f, // Starts with 'S' or '3'
	PrivateKeyID:            0x80,
	WitnessPubKeyHashAddrID: 0x06,
	WitnessScriptHashAddrID: 0x0A,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x88, 0xad, 0xe4}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x04, 0x88, 0xb2, 0x1e}, // starts with xpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 20, // SLIP-0044 registered coin type for DigiByte
}

// DigiByteTestNetParams defines the network parameters for the DigiByte test
// network.
var DigiByteTestNetParams = bitcoinCfg.Params{
	Name:        "digibyte-testnet",
	Net:         wire.BitcoinNet(0xfcc1b7dc), // DigiByte testnet magic
	DefaultPort: "12026",

	// Human-readable part for Bech32 encoded segwit addresses.
	Bech32HRPSegwit: "dgbt",

	// Address encoding magics
	PubKeyHashAddrID:        0x7e, // starts with 'm' or 'n'
	ScriptHashAddrID:        0x8c, // starts with 't'
	PrivateKeyID:            0xfe,
	WitnessPubKeyHashAddrID: 0x06,
	WitnessScriptHashAddrID: 0x0A,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 1, // Testnet coin type
}

// DigiByteRegTestParams defines the network parameters for the DigiByte
// regression test network. This network is similar to mainnet but is intended
// for private development and testing.
var DigiByteRegTestParams = bitcoinCfg.Params{
	Name:        "digibyte-regtest",
	Net:         wire.BitcoinNet(0xdab5bffa), // DigiByte regtest magic
	DefaultPort: "18444",

	// Human-readable part for Bech32 encoded segwit addresses.
	Bech32HRPSegwit: "dgbrt",

	// Address encoding magics
	PubKeyHashAddrID:        0x7e,
	ScriptHashAddrID:        0x8c,
	PrivateKeyID:            0xfe,
	WitnessPubKeyHashAddrID: 0x06,
	WitnessScriptHashAddrID: 0x0A,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94},
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf},

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 1,
}

// Register DigiByte networks with btcd
func init() {
	// Register DigiByte mainnet
	mustRegister := func(params *bitcoinCfg.Params) {
		err := bitcoinCfg.Register(params)
		if err != nil {
			panic("failed to register DigiByte network: " + err.Error())
		}
	}

	mustRegister(&DigiByteMainNetParams)
	mustRegister(&DigiByteTestNetParams)
	mustRegister(&DigiByteRegTestParams)
}
