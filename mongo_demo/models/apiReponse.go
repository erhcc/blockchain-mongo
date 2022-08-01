package models
/*
{
"id":1,
"jsonrpc":"2.0",
"result": {
    "hash":"0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b",
    "nonce":"0x",
    "blockHash": "0xbeab0aa2411b7ab17f30a99d3cb9c6ef2fc5426d6ad6fd9e2a26a6aed1d1055b",
    "blockNumber": "0x15df", // 5599
    "transactionIndex":  "0x1", // 1
    "from":"0x407d73d8a49eeb85d32cf465507dd71d507100c1",
    "to":"0x85h43d8a49eeb85d32cf465507dd71d507100c1",
    "value":"0x7f110", // 520464
    "gas": "0x7f110", // 520464
    "gasPrice":"0x09184e72a000",
    "input":"0x603880600c6000396000f300603880600c6000396000f3603880600c6000396000f360",
  }
}
*/

//uint64 Nonce/Gas/GasPrice
type ApiTransaction struct{
	Hash     string `json:"hash"`
	Nonce    string `json:"nonce"`
	BlockHash     string `json:"blockHash"`
	BlockNumber     string `json:"blockNumber"`
	TransactionIndex     string `json:"transactionIndex"`
	From       string `json:"from"`
	To       string `json:"to"`
	Value    string `json:"value"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Input  string   `json:"input"`
}

/* eth_getTransactionByHash
*/
type ApiTransactiondataResponse struct{
	Id     int `json:"id"`
	Jsonrpc    string `json:"jsonrpc"`
	Result	ApiTransaction `json:"result"`
}


/*
    "number": "0x1b4", // 436
    "hash": "0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331",
    "parentHash": "0x9646252be9520f6e71339a8df9c55e4d7619deeb018d2a3f2d21fc165dde5eb5",
    "nonce": "0xe04d296d2460cfb8472af2c5fd05b5a214109c25688d3704aed5484f9a7792f2",
    "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
    "logsBloom": "0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331",
    "transactionsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
    "stateRoot": "0xd5855eb08b3387c0af375e9cdb6acfc05eb8f519e419b874b6ff2ffda7ed1dff",
    "miner": "0x4e65fda2159562a496f9f3522f89122a3088497a",
    "difficulty": "0x027f07", // 163591
    "totalDifficulty":  "0x027f07", // 163591
    "extraData": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "size":  "0x027f07", // 163591
    "gasLimit": "0x9f759", // 653145
    "gasUsed": "0x9f759", // 653145
    "timestamp": "0x54e34e8e" // 1424182926
    "transactions": ["0xcdf05df923f6e418505750069d6486276b15fcc3cd2f42a7044c642d19a86d51",
      "0x0c66977ed87db75074cb2bea66b254af3b20bb3315e8095290ceb1260b1b7449",] 
    "uncles": ["0x1606e5...", "0xd5145a9..."]
*/

type ApiBlock struct{
	Number     string `number:"hash"`
	Hash     string `json:"hash"`
	ParentHash     string `json:"parentHash"`
	Nonce    string `json:"nonce"`
	Sha3Uncles    string `json:"sha3Uncles"`
	LogsBloom    string `json:"logsBloom"`
	TransactionsRoot    string `json:"transactionsRoot"`
	MixHash    string `json:"mixHash"`
	ReceiptsRoot    string `json:"receiptsRoot"`
	StateRoot    string `json:"stateRoot"`
	Miner    string `json:"miner"`
	Difficulty    string `json:"difficulty"`
	TotalDifficulty    string `json:"totalDifficulty"`
	ExtraData    string `json:"extraData"`
	Size    string `json:"size"`
	GasLimit string `json:"gasLimit"`
	GasUsed string `json:"gasUsed"`
	Timestamp  string   `json:"timestamp"`
	Transactions     []string `json:"transactions"`
	Uncles     []string `json:"uncles"`
}

/* eth_getBlockByNumber
*/
type ApiBlockdataResponse struct{
	Id     int `json:"id"`
	Jsonrpc    string `json:"jsonrpc"`
	Result	ApiBlock `json:"result"`
}

