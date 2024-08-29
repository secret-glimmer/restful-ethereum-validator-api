package services

import (
	"encoding/json"
	"fmt"
	"math/big"
	"validator-api/models"
	f "validator-api/pkgs/functions"
)

func BlockByTimeStamp(slot int, quickNodeUrl string, response *models.ResponseBlock) error {
	time := f.TimeFromSlot(slot)
	hexOfTime := fmt.Sprintf("0x%x", time.Unix())
	request := models.RequestBlock{
		Method: "erigon_getBlockByTimestamp",
		Params: []any{
			hexOfTime,
			true,
		},
		Id:      1,
		JsonRpc: "2.0",
	}
	data, err := Post(quickNodeUrl, request)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, response)
	if err != nil {
		return err
	}

	return nil
}

func BurntFee(blockNumber int64, block *models.Block) *big.Float {
	if blockNumber < 12965001 {
		return new(big.Float).SetFloat64(0)
	}
	baseFee := f.FloatFromHex(block.BaseFeePerGas)
	gasUsed := f.FloatFromHex(block.GasUsed)
	return baseFee.Mul(baseFee, gasUsed)
}

func StaticBlockReward(blockNumber int64) *big.Float {
	if blockNumber < 4370000 {
		return new(big.Float).SetFloat64(5)
	} else if blockNumber < 7280000 {
		return new(big.Float).SetFloat64(3)
	} else if blockNumber < 15537393 {
		return new(big.Float).SetFloat64(2)
	}
	return new(big.Float).SetFloat64(0)
}
