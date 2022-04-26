package btcrpc

import (
	"encoding/json"
	"fmt"
	"strings"
)

const AddressRoute string = "/address/"

type AddressSummary struct {
	Base58 struct {
		Hash    string `json:"hash"`
		Version int    `json:"version"`
	} `json:"base58"`
	Encoding        string `json:"encoding"`
	ValidateAddress struct {
		IsValid      bool   `json:"isvalid"`
		Address      string `json:"address"`
		ScriptPubKey string `json:"scriptPubKey"`
		IsScript     bool   `json:"isscript"`
		IsWitness    bool   `json:"iswitness"`
	} `json:"validateaddress"`
	ElectrumScriptHash string `json:"electrumScriptHash"`
	TXHistory          struct {
		TXCount            int            `json:"txCount"`
		TXIDs              []string       `json:"txids"`
		BlockHeightsByTxid map[string]int `json:"blockHeightsByTxid"`
		BalanceSat         int            `json:"balanceSat"`
		Request            struct {
			Limit  int    `json:"limit"`
			Offset int    `json:"offset"`
			Sort   string `json:"sort"`
		} `json:"request"`
	} `json:"txHistory"`
}

// AddressSummary takes a Bitcoin address and returns an AddressSummary object
func (c Config) AddressSummary(address string) (AddressSummary, error) {
	var AddressSummary AddressSummary
	// Trim any trailing slash
	endpoint := strings.TrimRight(c.APIEndpoint, "/")
	body, err := callAPI(endpoint + api + AddressRoute + address)
	if err != nil {
		return AddressSummary, err
	}

	err = json.Unmarshal(body, &AddressSummary)
	if err != nil {
		return AddressSummary, fmt.Errorf("error unmarshalling json: %w, body: %v", err, string(body))
	}
	return AddressSummary, nil
}
