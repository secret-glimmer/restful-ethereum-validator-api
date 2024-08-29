package services

import (
	"encoding/json"
	"validator-api/models"
)

func BlockReceiptByNumber(blockNumber string, quickNodeUrl string, response *models.ResponseBlockReceipt) error {
	request := models.RequestBlock{
		Method: "eth_getBlockReceipts",
		Params: []any{
			blockNumber,
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
