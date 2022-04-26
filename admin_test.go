package btcapi

import "testing"

func TestVersion(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.Version()
	if err != nil {
		t.Error(err)
	}
}
