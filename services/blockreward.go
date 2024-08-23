package services

import (
	"encoding/json"
	"strconv"
	"validator-api/models"
	cfg "validator-api/pkgs/config"

	"github.com/go-resty/resty/v2"
)

func GetBlockReward(slot int, config *cfg.Config) (models.BlockReward, bool, bool) {
	quicknodeHTTP := config.QuckNode.Http

	client := resty.New()
	resp, err := client.R().Get(quicknodeHTTP + "/blockreward/" + strconv.Itoa(slot))

	reward := models.BlockReward{}

	if err != nil {
		return reward, false, true
	}

	if resp.StatusCode() == 404 {
		return reward, true, false
	} else if resp.StatusCode() != 200 {
		return reward, false, true
	}

	err = json.Unmarshal(resp.Body(), &reward)

	if err != nil {
		return reward, false, true
	}

	return reward, false, false
}
