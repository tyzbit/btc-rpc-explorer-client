package btcapi

import "testing"

func TestBlockWithHash(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.BlockWithHash("0000000000000000001c8018d9cb3b742ef25114f27563e3fc4a1902167f9893")
	if err != nil {
		t.Error(err)
	}
}

func TestBlockWithHeight(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.BlockWithHeight(123456)
	if err != nil {
		t.Error(err)
	}
}
