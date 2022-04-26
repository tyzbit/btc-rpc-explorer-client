package btcrpc

import "testing"

func TestAddressSummary(t *testing.T) {
	btcrpc := Config{
		APIEndpoint: TestingEndpoint,
	}
	_, err := btcrpc.AddressSummary("34rng4QwB5pHUbGDJw1JxjLwgEU8TQuEqv")
	if err != nil {
		t.Error(err)
	}
}
