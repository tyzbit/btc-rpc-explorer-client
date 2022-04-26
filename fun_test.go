package btcapi

import "testing"

func TestQuotes(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.Quotes()
	if err != nil {
		t.Error(err)
	}
}

func TestQuote(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.Quote(0)
	if err != nil {
		t.Error(err)
	}
}

func TestRandomQuote(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.RandomQuote()
	if err != nil {
		t.Error(err)
	}
}
