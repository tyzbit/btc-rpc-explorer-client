package btcapi

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
)

type MiningHashrate struct {
	OneDay    HashrateSummary `json:"1Day"`
	SevenDay  HashrateSummary `json:"7Day"`
	ThirtyDay HashrateSummary `json:"30Day"`
	NinetyDay HashrateSummary `json:"90Day"`
	OneYear   HashrateSummary `json:"365Day"`
}

type HashrateSummary struct {
	Val              float64 `json:"val"`
	Unit             string  `json:"unit"`
	UnitAbbreviation string  `json:"unitAbbreviation"`
	UnitExponent     string  `json:"unitExponent"`
	UnitMultiplier   big.Int `json:"unitMultiplier"`
	Raw              big.Int `json:"raw"`
	String1          string  `json:"string1"`
	String2          string  `json:"string2"`
	String3          string  `json:"string3"`
}

type BlockMiningDetails struct {
	TXCount    int     `json:"txCount"`
	MinFeeRate float64 `json:"minFeeRate"`
	MaxFeeRate float64 `json:"maxFeeRate"`
	MinFeeTXID string  `json:"minFeeTxid"`
	MaxFeeTXID string  `json:"maxFeeTxid"`
	TotalFees  float64 `json:"totalFees"`
}

type Included struct {
	Included bool `json:"included"`
}

type MinerSummary struct {
	Miners map[string]struct {
		Name    string `json:"name"`
		Details struct {
			Name         string `json:"name"`
			Link         string `json:"link"`
			IdentifiedBy string `json:"identifiedBy"`
		} `json:"details"`
		Blocks            []int  `json:"blocks"`
		TotalFees         string `json:"totalFees"`
		TotalSubsidy      string `json:"totalSubsidy"`
		TotalTransactions int    `json:"totalTransactions"`
		TotalWeight       int    `json:"totalWeight"`
		SubsidyCount      int    `json:"subsidyCount"`
	} `json:"miners"`
	MinerNames []string `json:"minerNamesSortedByBlockCount"` // Sorted by block count
	Overall    struct {
		BlockCount        int    `json:"blockCount"`
		TotalFees         string `json:"totalFees"`
		TotalSubsidy      string `json:"totalSubsidy"`
		TotalTransactions int    `json:"totalTransactions"`
		TotalWeight       int    `json:"totalWeight"`
		SubsidyCount      int    `json:"subsidyCount"`
	} `json:"overall"`
}

const MiningRoute string = "/mining"

// HashRate returns details of the current hashrate.
func (c Config) Hashrate() (hashrate MiningHashrate, err error) {
	body, err := getAPI(c.ExplorerURL + api + MiningRoute + "/hashrate")
	if err != nil {
		return hashrate, err
	}

	err = json.Unmarshal(body, &hashrate)
	if err != nil {
		return hashrate, fmt.Errorf("unable to parse returned body: %v, err: %w", string(body), err)
	}
	return hashrate, nil
}

// UTXOSet returns the current UTXO snapshot.
func (c Config) DifficultyAdjustmentEstimate() (estimate float64, err error) {
	body, err := getAPI(c.ExplorerURL + api + MiningRoute + "/diff-adj-estimate")
	if err != nil {
		return estimate, err
	}

	estimate, err = strconv.ParseFloat(string(body), 64)
	if err != nil {
		return estimate, fmt.Errorf("unable to parse returned body: %v, err: %w", string(body), err)
	}
	return estimate, nil
}

// NextBlock returns details about the nextblock.
func (c Config) NextBlock() (details BlockMiningDetails, err error) {
	body, err := getAPI(c.ExplorerURL + api + MiningRoute + "/next-block")
	if err != nil {
		return details, err
	}

	err = json.Unmarshal(body, &details)
	if err != nil {
		return details, fmt.Errorf("unable to parse returned body: %v, err: %w", string(body), err)
	}
	return details, nil
}

// NextBlockTXIDs returns transaction IDs that might be included in the next block.
func (c Config) NextBlockTXIDs() (txids []string, err error) {
	body, err := getAPI(c.ExplorerURL + api + MiningRoute + "/next-block/txids")
	if err != nil {
		return txids, err
	}

	err = json.Unmarshal(body, &txids)
	if err != nil {
		return txids, fmt.Errorf("unable to parse returned body: %v, err: %w", string(body), err)
	}
	return txids, nil
}

// NextBlockIncludes returns whether a transaction is likely to be included in the next block.
func (c Config) NextBlockIncludes(txid string) (included bool, err error) {
	body, err := getAPI(c.ExplorerURL + api + MiningRoute + "/next-block/includes/" + txid)
	if err != nil {
		return included, err
	}

	var response Included
	err = json.Unmarshal(body, &response)
	if err != nil {
		return included, fmt.Errorf("unable to parse returned body: %v, err: %w", string(body), err)
	}
	included = response.Included
	return included, nil
}

// MinerSummery returns miner summary since the period specified (ex: 1d)
func (c Config) MinerSummary(since string) (summary MinerSummary, err error) {
	body, err := getAPI(c.ExplorerURL + api + MiningRoute + "/miner-summary?since=" + since)
	if err != nil {
		return summary, err
	}

	err = json.Unmarshal(body, &summary)
	if err != nil {
		return summary, fmt.Errorf("unable to parse returned body: %v, err: %w", string(body), err)
	}
	return summary, nil
}
