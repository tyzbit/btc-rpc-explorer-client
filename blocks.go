package btcrpc

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const BlockRoute string = "/block/"
const BlockTipRoute string = "/blocks/tip"

type BlockDetails struct {
	Hash              string   `json:"hash"`
	Confirmations     int      `json:"confirmations"`
	Height            int      `json:"Height"`
	Version           int      `json:"version"`
	VersionHex        string   `json:"versionHex"`
	MerkleRoot        string   `json:"merkleroot"`
	Time              int      `json:"time"`
	MedianTime        int      `json:"mediantime"`
	Nonce             int      `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        string   `json:"difficulty"`
	ChainWork         string   `json:"chainwork"`
	NumTx             int      `json:"nTx"`
	PreviousBlockHash string   `json:"previousblockhash"`
	NextBlockHash     string   `json:"nextblockhash"`
	StrippedSize      int      `json:"strippedsize"`
	Size              int      `json:"size"`
	Weight            int      `json:"weight"`
	Tx                []string `json:"tx"`
	CoinbaseTx        struct {
		InActiveChain bool   `json:"in_active_chain"`
		TxID          string `json:"txid"`
		Hash          string `json:"hash"`
		Version       int    `json:"version"`
		Size          int    `json:"size"`
		VSize         int    `json:"vsize"`
		Weight        int    `json:"weight"`
		LockTime      int    `json:"locktime"`
		VIn           []struct {
			Coinbase    string   `json:"coinbase"`
			TxInWitness []string `json:"txinwitness"`
			Sequence    int      `json:"sequence"`
		} `json:"vin"`
		VOut []struct {
			Value        float64 `json:"value"`
			N            int     `json:"n"`
			ScriptPubKey struct {
				ASM     string `json:"asm"`
				Hex     string `json:"hex"`
				Address string `json:"address"`
				Type    string `json:"type"`
			} `json:"scriptPubKey"`
		} `json:"vout"`
		Hex           string `json:"hex"`
		BlockHash     string `json:"blockhash"`
		Confirmations int    `json:"confirmations"`
		Time          int    `json:"time"`
		BlockTime     int    `json:"blocktime"`
		TotalFees     string `json:"totalFees"`
		Miner         struct {
			Name         string `json:"name"`
			Link         string `json:"link"`
			IdentifiedBy string `json:"identifiedBy"`
		} `json:"miner"`
		Subsidy string `json:"subsidy"`
	} `json:"coinbaseTx"`
}

// Block takes block hash and returns a BlockDetails object
func (c Config) BlockWithHash(hash string) (BlockDetails, error) {
	var BlockDetails BlockDetails
	// Trim any trailing slash
	endpoint := strings.TrimRight(c.APIEndpoint, "/")
	body, err := callAPI(endpoint + api + BlockRoute + hash)
	if err != nil {
		return BlockDetails, err
	}

	err = json.Unmarshal(body, &BlockDetails)
	if err != nil {
		return BlockDetails, fmt.Errorf("error unmarshalling json: %w, body: %v", err, string(body))
	}
	return BlockDetails, nil
}

// Block takes block height and returns a BlockDetails object
func (c Config) BlockWithHeight(height int) (BlockDetails, error) {
	var BlockDetails BlockDetails
	// Trim any trailing slash
	endpoint := strings.TrimRight(c.APIEndpoint, "/")
	body, err := callAPI(endpoint + api + BlockRoute + fmt.Sprintf("%d", height))
	if err != nil {
		return BlockDetails, err
	}

	err = json.Unmarshal(body, &BlockDetails)
	if err != nil {
		return BlockDetails, fmt.Errorf("error unmarshalling json: %w, body: %v", err, string(body))
	}
	return BlockDetails, nil
}

// TipHeight returns the current tip height.
func (c Config) TipHeight() (int, error) {
	// Trim any trailing slash
	endpoint := strings.TrimRight(c.APIEndpoint, "/")
	body, err := callAPI(endpoint + api + BlockTipRoute + "/height")
	if err != nil {
		return 0, err
	}

	tip, err := strconv.Atoi(string(body))
	if err != nil {
		return 0, fmt.Errorf("unable to parse returned body: %v", string(body))
	}

	return tip, nil
}

// TipHash returns the tip hash.
func (c Config) TipHash() (string, error) {
	// Trim any trailing slash
	endpoint := strings.TrimRight(c.APIEndpoint, "/")
	body, err := callAPI(endpoint + api + BlockTipRoute + "/hash")
	if err != nil {
		return "", err
	}

	return string(body), nil
}
