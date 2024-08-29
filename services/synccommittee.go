package services

import (
	"encoding/json"
	"fmt"
	"validator-api/models"
)

func SyncCommitteeIndices(slot int, quickNodeUrl string, keys *[]string) error {
	data, err := Post(quickNodeUrl+fmt.Sprintf("/eth/v1/beacon/rewards/sync_committee/%d", slot), []string{})
	if err != nil {
		return err
	}

	response := models.ResponseValidator{}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	for _, validator := range response.Data {
		*keys = append(*keys, validator.ValidatorIndex)
	}

	return nil
}
