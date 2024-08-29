package services

import (
	"math/big"
	"validator-api/models"
	f "validator-api/pkgs/functions"
)

func TransactionFee(block models.Block, quickNodeUrl string) *big.Float {
	sum := new(big.Float).SetFloat64(0)

	response := models.ResponseBlockReceipt{}

	blockGasUsed := map[string]string{}

	BlockReceiptByNumber(block.Number, quickNodeUrl, &response)

	for _, receipt := range response.Result {
		blockGasUsed[receipt.TransactionHash] = receipt.GasUsed
	}

	for _, transaction := range block.Transactions {
		gasUsed := f.FloatFromHex(blockGasUsed[transaction.Hash])
		gasPrice := f.FloatFromHex(transaction.GasPrice)
		gasFee := gasUsed.Mul(gasUsed, gasPrice)
		sum = sum.Add(sum, gasFee)
	}

	return sum
}
