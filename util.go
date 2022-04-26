package btcapi

import (
	"encoding/json"
	"fmt"
)

type ExtendedPublicKeyDetails struct {
	KeyType          string       `json:"keyType"`
	OutputType       string       `json:"outputType"`
	OutputTypeDesc   string       `json:"outputTypeDesc"`
	BIP32Path        string       `json:"bip32Path"`
	RelatedKeys      []RelatedKey `json:"relatedKeys"`
	ReceiveAddresses []string     `json:"receiveAddresses"`
	ChangeAddresses  []string     `json:"changeAddresses"`
}

type RelatedKey struct {
	KeyType      string `json:"keyType"`
	Key          string `json:"key"`
	OutputType   string `json:"outputType"`
	FirstAddress string `json:"firstAddress"`
}

const UtilRoute string = "/util"

// ExtendedPublicKeyDetails returns details for an extended public key.
func (c Config) ExtendedPublicKeyDetails(pubkey string) (details ExtendedPublicKeyDetails, err error) {
	body, err := getAPI(c.ExplorerURL + api + UtilRoute + "/xyzpub/" + pubkey)
	if err != nil {
		return details, err
	}

	err = json.Unmarshal(body, &details)
	if err != nil {
		return details, fmt.Errorf("unable to parse returned body: %v, err: %w", string(body), err)
	}
	return details, nil
}
