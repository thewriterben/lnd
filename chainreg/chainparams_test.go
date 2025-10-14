package chainreg

import (
	"testing"

	"github.com/lightningnetwork/lnd/chainparams"
	"github.com/lightningnetwork/lnd/keychain"
)

// TestDigiByteMainNetParams verifies the DigiByte mainnet chain registry
// parameters.
func TestDigiByteMainNetParams(t *testing.T) {
	// Verify params are set correctly
	if DigiByteMainNetParams.Params != &chainparams.DigiByteMainNetParams {
		t.Error("DigiByteMainNetParams.Params not set correctly")
	}

	// Verify RPC port
	if DigiByteMainNetParams.RPCPort != "14022" {
		t.Errorf("Expected RPC port 14022, got %s",
			DigiByteMainNetParams.RPCPort)
	}

	// Verify coin type
	if DigiByteMainNetParams.CoinType != keychain.CoinTypeDigiByte {
		t.Errorf("Expected coin type %d, got %d",
			keychain.CoinTypeDigiByte, DigiByteMainNetParams.CoinType)
	}
}

// TestDigiByteTestNetParams verifies the DigiByte testnet chain registry
// parameters.
func TestDigiByteTestNetParams(t *testing.T) {
	// Verify params are set correctly
	if DigiByteTestNetParams.Params != &chainparams.DigiByteTestNetParams {
		t.Error("DigiByteTestNetParams.Params not set correctly")
	}

	// Verify RPC port
	if DigiByteTestNetParams.RPCPort != "14023" {
		t.Errorf("Expected RPC port 14023, got %s",
			DigiByteTestNetParams.RPCPort)
	}

	// Verify coin type (testnet uses general testnet coin type)
	if DigiByteTestNetParams.CoinType != keychain.CoinTypeTestnet {
		t.Errorf("Expected coin type %d, got %d",
			keychain.CoinTypeTestnet, DigiByteTestNetParams.CoinType)
	}
}

// TestDigiByteRegTestParams verifies the DigiByte regtest chain registry
// parameters.
func TestDigiByteRegTestParams(t *testing.T) {
	// Verify params are set correctly
	if DigiByteRegTestParams.Params != &chainparams.DigiByteRegTestParams {
		t.Error("DigiByteRegTestParams.Params not set correctly")
	}

	// Verify RPC port
	if DigiByteRegTestParams.RPCPort != "18443" {
		t.Errorf("Expected RPC port 18443, got %s",
			DigiByteRegTestParams.RPCPort)
	}

	// Verify coin type (regtest uses general testnet coin type)
	if DigiByteRegTestParams.CoinType != keychain.CoinTypeTestnet {
		t.Errorf("Expected coin type %d, got %d",
			keychain.CoinTypeTestnet, DigiByteRegTestParams.CoinType)
	}
}

// TestDigiByteCoinTypes verifies the coin types are set correctly.
func TestDigiByteCoinTypes(t *testing.T) {
	// Mainnet should use DigiByte coin type (20)
	if DigiByteMainNetParams.CoinType != 20 {
		t.Errorf("DigiByte mainnet should use coin type 20, got %d",
			DigiByteMainNetParams.CoinType)
	}

	// Testnet and regtest should use testnet coin type (1)
	if DigiByteTestNetParams.CoinType != 1 {
		t.Errorf("DigiByte testnet should use coin type 1, got %d",
			DigiByteTestNetParams.CoinType)
	}

	if DigiByteRegTestParams.CoinType != 1 {
		t.Errorf("DigiByte regtest should use coin type 1, got %d",
			DigiByteRegTestParams.CoinType)
	}
}

// TestDigiByteDefaults verifies DigiByte-specific default values.
func TestDigiByteDefaults(t *testing.T) {
	// Verify DigiByte timelock delta is lower than Bitcoin's
	if DefaultDigiByteTimeLockDelta >= DefaultBitcoinTimeLockDelta {
		t.Errorf("DigiByte timelock delta (%d) should be lower than Bitcoin's (%d)",
			DefaultDigiByteTimeLockDelta, DefaultBitcoinTimeLockDelta)
	}

	// Verify DigiByte base fee is lower than Bitcoin's
	if DefaultDigiByteBaseFeeMSat >= DefaultBitcoinBaseFeeMSat {
		t.Errorf("DigiByte base fee (%d) should be lower than Bitcoin's (%d)",
			DefaultDigiByteBaseFeeMSat, DefaultBitcoinBaseFeeMSat)
	}

	// Verify DigiByte static fee is lower than Bitcoin's
	if DefaultDigiByteStaticFeePerKW >= DefaultBitcoinStaticFeePerKW {
		t.Errorf("DigiByte static fee (%d) should be lower than Bitcoin's (%d)",
			DefaultDigiByteStaticFeePerKW, DefaultBitcoinStaticFeePerKW)
	}
}
