package chainparams

import (
	"testing"

	"github.com/btcsuite/btcd/wire"
)

// TestDigiByteMainNetParams verifies the DigiByte mainnet parameters.
func TestDigiByteMainNetParams(t *testing.T) {
	// Verify network name
	if DigiByteMainNetParams.Name != "digibyte-mainnet" {
		t.Errorf("Expected name 'digibyte-mainnet', got '%s'",
			DigiByteMainNetParams.Name)
	}

	// Verify network magic
	expectedNet := wire.BitcoinNet(0xfac3b6da)
	if DigiByteMainNetParams.Net != expectedNet {
		t.Errorf("Expected net 0xfac3b6da, got %v",
			DigiByteMainNetParams.Net)
	}

	// Verify default port
	if DigiByteMainNetParams.DefaultPort != "12024" {
		t.Errorf("Expected port 12024, got %s",
			DigiByteMainNetParams.DefaultPort)
	}

	// Verify address prefixes
	if DigiByteMainNetParams.PubKeyHashAddrID != 0x1e {
		t.Errorf("Expected PubKeyHashAddrID 0x1e, got 0x%x",
			DigiByteMainNetParams.PubKeyHashAddrID)
	}

	if DigiByteMainNetParams.ScriptHashAddrID != 0x3f {
		t.Errorf("Expected ScriptHashAddrID 0x3f, got 0x%x",
			DigiByteMainNetParams.ScriptHashAddrID)
	}

	// Verify Bech32 HRP
	if DigiByteMainNetParams.Bech32HRPSegwit != "dgb" {
		t.Errorf("Expected Bech32HRP 'dgb', got '%s'",
			DigiByteMainNetParams.Bech32HRPSegwit)
	}

	// Verify coin type
	if DigiByteMainNetParams.HDCoinType != 20 {
		t.Errorf("Expected coin type 20, got %d",
			DigiByteMainNetParams.HDCoinType)
	}
}

// TestDigiByteTestNetParams verifies the DigiByte testnet parameters.
func TestDigiByteTestNetParams(t *testing.T) {
	// Verify network name
	if DigiByteTestNetParams.Name != "digibyte-testnet" {
		t.Errorf("Expected name 'digibyte-testnet', got '%s'",
			DigiByteTestNetParams.Name)
	}

	// Verify network magic
	expectedNet := wire.BitcoinNet(0xfcc1b7dc)
	if DigiByteTestNetParams.Net != expectedNet {
		t.Errorf("Expected net 0xfcc1b7dc, got %v",
			DigiByteTestNetParams.Net)
	}

	// Verify default port
	if DigiByteTestNetParams.DefaultPort != "12026" {
		t.Errorf("Expected port 12026, got %s",
			DigiByteTestNetParams.DefaultPort)
	}

	// Verify Bech32 HRP
	if DigiByteTestNetParams.Bech32HRPSegwit != "dgbt" {
		t.Errorf("Expected Bech32HRP 'dgbt', got '%s'",
			DigiByteTestNetParams.Bech32HRPSegwit)
	}

	// Verify coin type (testnet uses coin type 1)
	if DigiByteTestNetParams.HDCoinType != 1 {
		t.Errorf("Expected coin type 1, got %d",
			DigiByteTestNetParams.HDCoinType)
	}
}

// TestDigiByteRegTestParams verifies the DigiByte regtest parameters.
func TestDigiByteRegTestParams(t *testing.T) {
	// Verify network name
	if DigiByteRegTestParams.Name != "digibyte-regtest" {
		t.Errorf("Expected name 'digibyte-regtest', got '%s'",
			DigiByteRegTestParams.Name)
	}

	// Verify network magic
	expectedNet := wire.BitcoinNet(0xdab5bffa)
	if DigiByteRegTestParams.Net != expectedNet {
		t.Errorf("Expected net 0xdab5bffa, got %v",
			DigiByteRegTestParams.Net)
	}

	// Verify default port
	if DigiByteRegTestParams.DefaultPort != "18444" {
		t.Errorf("Expected port 18444, got %s",
			DigiByteRegTestParams.DefaultPort)
	}

	// Verify Bech32 HRP
	if DigiByteRegTestParams.Bech32HRPSegwit != "dgbrt" {
		t.Errorf("Expected Bech32HRP 'dgbrt', got '%s'",
			DigiByteRegTestParams.Bech32HRPSegwit)
	}
}

// TestDigiByteNetworkMagics verifies that DigiByte networks have unique magic
// numbers.
func TestDigiByteNetworkMagics(t *testing.T) {
	mainnetMagic := DigiByteMainNetParams.Net
	testnetMagic := DigiByteTestNetParams.Net
	regtestMagic := DigiByteRegTestParams.Net

	// Ensure all network magics are unique
	if mainnetMagic == testnetMagic {
		t.Error("Mainnet and testnet have the same magic number")
	}
	if mainnetMagic == regtestMagic {
		t.Error("Mainnet and regtest have the same magic number")
	}
	if testnetMagic == regtestMagic {
		t.Error("Testnet and regtest have the same magic number")
	}
}

// TestDigiByteAddressPrefixes verifies the address prefixes are correct.
func TestDigiByteAddressPrefixes(t *testing.T) {
	// Mainnet P2PKH addresses should start with 'D' (0x1e)
	// This can be verified by knowing that:
	// - 'D' in base58 corresponds to a version byte of 0x1e
	if DigiByteMainNetParams.PubKeyHashAddrID != 0x1e {
		t.Errorf("Mainnet P2PKH prefix should be 0x1e for 'D' addresses, got 0x%x",
			DigiByteMainNetParams.PubKeyHashAddrID)
	}

	// Testnet P2PKH addresses should start with 'm' or 'n' (0x6f or 0x7e)
	// DigiByte testnet uses 0x7e
	if DigiByteTestNetParams.PubKeyHashAddrID != 0x7e {
		t.Errorf("Testnet P2PKH prefix should be 0x7e, got 0x%x",
			DigiByteTestNetParams.PubKeyHashAddrID)
	}
}
