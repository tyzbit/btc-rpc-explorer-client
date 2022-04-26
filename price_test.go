package btcapi

import "testing"

func TestPrice(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.Price()
	if err != nil {
		t.Error(err)
	}
}

func TestPriceIn(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	for _, currency := range []string{"usd", "eur", "gbp", "xau"} {
		_, err := btcapi.PriceIn(currency)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestMarketCapIn(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	for _, currency := range []string{"usd", "eur", "gbp", "xau"} {
		_, err := btcapi.MarketCapIn(currency)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestPriceInSats(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	for _, currency := range []string{"usd", "eur", "gbp", "xau"} {
		_, err := btcapi.PriceInSats(currency)
		if err != nil {
			t.Error(err)
		}
	}
}
