package config


type APIConfig struct {
	URL               string `json:"url"`
	Method            string `json:"method"`
	RequestStructure  string `json:"request_structure"`
	ResponseStructure string `json:"response_structure"`
}

type TransactionAPI struct {
	APIIndex   int               `json:"api_index"`
	Sequence   int               `json:"sequence"`
	Dependency map[string]string `json:"dependency"`
}

type Transaction struct {
	Name string          `json:"name"`
	APIs []TransactionAPI `json:"apis"`
}

type ConfigFile struct {
	APIs         []APIConfig  `json:"apis"`
	Transactions []Transaction `json:"transactions"`
}