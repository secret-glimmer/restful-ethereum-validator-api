package services

import (
	"encoding/json"
	"strconv"
	cfg "validator-api/pkgs/config"

	"github.com/go-resty/resty/v2"
)

func GetSyncDuties(slot int, config *cfg.Config) ([]string, bool, bool) {
	quicknodeHTTP := config.QuckNode.Http

	client := resty.New()
	resp, err := client.R().Get(quicknodeHTTP + "/syncduties/" + strconv.Itoa(slot))

	if err != nil {
		return nil, false, true
	}

	if resp.StatusCode() == 404 {
		return nil, true, false
	} else if resp.StatusCode() != 200 {
		return nil, false, true
	}

	keys := []string{}

	json.Unmarshal(resp.Body(), &keys)

	return keys, false, false
}
