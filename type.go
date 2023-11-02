package trongrid

type Error struct {
	Error string `json:"error"`
}

type Meta struct {
	Links       *MetaLinks `json:"links"`
	Fingerprint string     `json:"fingerprint"`
	At          int64      `json:"at"`
	PageSize    int32      `json:"page_size"`
}

type MetaLinks struct {
	Next string `json:"next"`
}

type Token struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int32  `json:"decimals"`
}

type Transaction struct {
	Token          *Token          `json:"token_info"`
	From           string          `json:"from"`
	ID             string          `json:"transaction_id"`
	To             string          `json:"to"`
	Type           TransactionType `json:"type"`
	Value          string          `json:"value"`
	BlockTimestamp int64           `json:"block_timestamp"`
}

type TransactionType string
