package storage

type Opcode struct {
	BlockNumber int64 `json:”blockNumber,omitempty”`
	Sender string `json:”sender,omitempty”`
	ContractAddress string `json:”contractAddress,omitempty”`
	Opcode string `json:”opcode,omitempty”`
	ElapsedTime int64 `json:”elapsedTime,omitempty”`
	WorkerId int
}

type Transaction struct {
	ElapsedTime int64 `json:”elapsedTime,omitempty”`
	GasUsed int64 `json:”gasUsed,omitempty”`
	GasPrice string `json:"gasPrice,omitempty"`
	GasAmount int64 `json:"gasAmount,omitempty"`
	Data string `json:"data,omitempty"`
	Sender string `json:"sender,omitempty"`
	ToAddress string `json:"sender,omitempty"`
	Nonce int64 `json:"nonce,omitempty"`
	BlockNumber int64 `json:"nonce,omitempty""`
}