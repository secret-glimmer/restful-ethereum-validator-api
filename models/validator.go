package models

type Validator struct {
	ValidatorIndex string `json:"validator_index"`
	Reward         string `json:"reward"`
}

type ResponseValidator struct {
	ExecutionOptimistic bool        `json:"execution_optimistic"`
	Finalized           bool        `json:"finalized"`
	Data                []Validator `json:"data"`
}
