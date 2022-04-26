package btcapi

import "testing"

func TestCoins(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.Coins()
	if err != nil {
		t.Error(err)
	}
}

func TestUTXOSet(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.UTXOSet()
	if err != nil {
		t.Error(err)
	}
}
