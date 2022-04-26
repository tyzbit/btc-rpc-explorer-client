package btcapi

import "testing"

func TestHashrate(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.Hashrate()
	if err != nil {
		t.Error(err)
	}
}

func TestDifficultyAdjustmentEstimate(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.DifficultyAdjustmentEstimate()
	if err != nil {
		t.Error(err)
	}
}

func TestNextBlock(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.NextBlock()
	if err != nil {
		t.Error(err)
	}
}

func TestNextBlockTXIDs(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.NextBlockTXIDs()
	if err != nil {
		t.Error(err)
	}
}

func TestNextBlockIncludes(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.NextBlockIncludes("yourTxId")
	if err != nil {
		t.Error(err)
	}
}

func TestMinerSummary(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.MinerSummary("1d")
	if err != nil {
		t.Error(err)
	}
}
