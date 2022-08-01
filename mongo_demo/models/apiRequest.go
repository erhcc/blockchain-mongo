package models

/*
{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params": ["0x5BAD55",false],"id":1}
*/
type ApiBlockRequest struct{
	Jsonrpc string  	`json:"jsonrpc"`
	Method string  	`json:"method"`
	Params []interface{}	`json:"params"`
	Id int  		`json:"id"`
}
/*
'{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params": ["0xbb3a336e3f823ec18197f1e13ee875700f08f03e2cab75f0d0b118dabb44cba0"],"id":1}'
*/
type ApiTransactionRequest struct{
	Jsonrpc string  	`json:"jsonrpc"`
	Method string  	`json:"method"`
	Params []string	`json:"params"`
	Id int  		`json:"id"`
}
