package btcapi

import "testing"

func TestTx(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.Tx("f4184fc596403b9d638783cf57adfe4c75c605f6356fbc91338530e9831e9e16")
	if err != nil {
		t.Error(err)
	}
}
