package btcapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Price struct {
	USD string `json:"usd"`
	EUR string `json:"eur"`
	GBP string `json:"gbp"`
	XAU string `json:"xau"`
}

const PriceRoute string = "/price"

// Price returns the price of 1BTC in USD, EUR, GBP and XAU.
func (c Config) Price() (price Price, err error) {
	body, err := getAPI(c.ExplorerURL + api + PriceRoute)
	if err != nil {
		return price, err
	}

	err = json.Unmarshal(body, &price)
	if err != nil {
		return price, fmt.Errorf("unable to parse returned body: %v, err: %w", string(body), err)
	}
	return price, nil
}

// Price returns the price of 1BTC in one of: USD, EUR, GBP and XAU.
func (c Config) PriceIn(currency string) (price string, err error) {
	body, err := getAPI(c.ExplorerURL + api + PriceRoute + "/" + strings.ToLower(currency))
	if err != nil {
		return price, err
	}

	price = string(body)
	if err != nil {
		return price, fmt.Errorf("unable to parse returned body: %v, err: %w", string(body), err)
	}
	return price, nil
}

// MarketCapIn returns the market cap in one of: USD, EUR, GBP and XAU.
func (c Config) MarketCapIn(currency string) (cap float64, err error) {
	body, err := getAPI(c.ExplorerURL + api + PriceRoute + "/" + strings.ToLower(currency) + "/marketcap")
	if err != nil {
		return cap, err
	}

	cap, err = strconv.ParseFloat(string(body), 64)
	if err != nil {
		return cap, fmt.Errorf("unable to parse returned body: %v, err: %w", string(body), err)
	}
	return cap, nil
}

// PriceInSats returns the value in satoshis of one unit of currency for one of:
// USD, EUR, GBP and XAU.
func (c Config) PriceInSats(currency string) (satoshis int, err error) {
	body, err := getAPI(c.ExplorerURL + api + PriceRoute + "/" + strings.ToLower(currency) + "/sats")
	if err != nil {
		return satoshis, err
	}

	satoshis, err = strconv.Atoi(string(body))
	if err != nil {
		return satoshis, fmt.Errorf("unable to parse returned body: %v, err: %w", string(body), err)
	}
	return satoshis, nil
}
