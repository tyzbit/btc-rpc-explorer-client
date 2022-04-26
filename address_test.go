package btcapi

import "testing"

func TestAddressSummary(t *testing.T) {
	btcapi := Config{
		APIEndpoint: TestingEndpoint,
	}
	_, err := btcapi.AddressSummary("34rng4QwB5pHUbGDJw1JxjLwgEU8TQuEqv")
	if err != nil {
		t.Error(err)
	}
}
