package btcapi

import "testing"

func TestMempoolCount(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.MempoolCount()
	if err != nil {
		t.Error(err)
	}
}

func TestMempoolFees(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.Fees()
	if err != nil {
		t.Error(err)
	}
}
