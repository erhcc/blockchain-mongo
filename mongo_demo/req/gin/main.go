package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/erichuang-code/blockchain-mongo/models"
	"github.com/gin-gonic/gin"
)

// type blockChainResponse struct{
// 	Result models.Transaction
// }

var apiTransactiondataResponse =models.ApiTransactiondataResponse{
	Id:      1,
	Jsonrpc: "2.0",
	Result:  models.ApiTransaction{
		Hash:"0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b",
		Nonce:"0x",
		BlockHash: "0xbeab0aa2411b7ab17f30a99d3cb9c6ef2fc5426d6ad6fd9e2a26a6aed1d1055b",
		BlockNumber: "0x15df", // 5599
		TransactionIndex:  "0x1", // 1
		From:"0x407d73d8a49eeb85d32cf465507dd71d507100c1",
		To:"0x85h43d8a49eeb85d32cf465507dd71d507100c1",
		Value:"0x7f110", // 520464
		Gas: "0x7f110", // 520464
		GasPrice:"0x09184e72a000",
		Input:"0x603880600c6000396000f300603880600c6000396000f3603880600c6000396000f360",
	},
}
//mock data returned by api
var apiBlockdataResponse =models.ApiBlockdataResponse{
	Id:      1,
	Jsonrpc: "2.0",
	Result:  models.ApiBlock{
		Difficulty: "0xbfabcdbd93dda",
		ExtraData: "0x737061726b706f6f6c2d636e2d6e6f64652d3132",
		GasLimit: "0x79f39e",
		GasUsed: "0x79ccd3",
		Hash: "0xb3b20624f8f0f86eb50dd04688409e5cea4bd02d700bf6e79e9384d47d6a5a35",
		LogsBloom: "0x4848112002a2020aaa0812180045840210020005281600c80104264300080008000491220144461026015300100000128005018401002090a824a4150015410020140400d808440106689b29d0280b1005200007480ca950b15b010908814e01911000054202a020b05880b914642a0000300003010044044082075290283516be82504082003008c4d8d14462a8800c2990c88002a030140180036c220205201860402001014040180002006860810ec0a1100a14144148408118608200060461821802c081000042d0810104a8004510020211c088200420822a082040e10104c00d010064004c122692020c408a1aa2348020445403814002c800888208b1",
		Miner: "0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c",
		MixHash: "0x3d1fdd16f15aeab72e7db1013b9f034ee33641d92f71c0736beab4e67d34c7a7",
		Nonce: "0x4db7a1c01d8a8072",
		Number: "0x5bad55",
		ParentHash: "0x61a8ad530a8a43e3583f8ec163f773ad370329b2375d66433eb82f005e1d6202",
		ReceiptsRoot: "0x5eced534b3d84d3d732ddbc714f5fd51d98a941b28182b6efe6df3a0fe90004b",
		Sha3Uncles: "0x8a562e7634774d3e3a36698ac4915e37fc84a2cd0044cb84fa5d80263d2af4f6",
		Size: "0x41c7",
		StateRoot: "0xf5208fffa2ba5a3f3a2f64ebd5ca3d098978bedd75f335f56b705d8715ee2305",
		Timestamp: "0x5b541449",
		TotalDifficulty: "0x12ac11391a2f3872fcd",
		Transactions: []string{
			"0x8784d99762bccd03b2086eabccee0d77f14d05463281e121a62abfebcf0d2d5f",
			"0x311be6a9b58748717ac0f70eb801d29973661aaf1365960d159e4ec4f4aa2d7f",
			"0xe42b0256058b7cad8a14b136a0364acda0b4c36f5b02dea7e69bfd82cef252a2",
		},
		TransactionsRoot: "0xf98631e290e88f58a46b7032f025969039aa9b5696498efc76baf436fa69b262",
		// uncles: []string{
		//   "0x824cce7c7c2ec6874b9fa9a9a898eb5f27cbaf3991dfa81084c3af60d1db618c",
		// }
	},
}

/* mock:
1. serve api to blockchain
	1.1 get blockdata
		请求参数：
		BLOCK PARAMETER：区块编号，或字符串标签 "latest"、"earliest" 、 "pending"
		FLAG ：是否返回完整交易对象，必需
		'{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params": ["0x5BAD55",false],"id":1}'
	1.2 get transactiondata within that block
		请求参数：
		TRANSACTION HASH ：要查询的交易的哈希，必需
		'{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params": ["0xbb3a336e3f823ec18197f1e13ee875700f08f03e2cab75f0d0b118dabb44cba0"],"id":1}'
*/
func main(){

	r:=gin.Default()

	r.POST("/blockdata",func (c *gin.Context)  {

		//get query param
		//jsonData, err := ioutil.ReadAll(c.Request.Body)
		var json models.ApiBlockRequest
		if err:=c.ShouldBindJSON(&json);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		}

		paramBlockNumber,ok:=json.Params[0].(string)//paramBlockNumber
		
		log.Printf("ApiBlockRequest is : %+v",json)
		log.Printf("paramBlockNumber is : %s",paramBlockNumber)

		var response models.ApiBlockdataResponse
		if ok {
			//send back the respone via paramBlockNumber, here a mock data
			response=apiBlockdataResponse	
			c.JSON(http.StatusOK,response)
		}else{
			c.JSON(http.StatusBadRequest,gin.H{"param incorrect":"block param incorrect"})
		}


	})

	r.POST("/transactiondata",func (c *gin.Context)  {
		var json models.ApiTransactionRequest
		if err:=c.ShouldBindJSON(&json);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		}

		paramTransaction:=json.Params

		log.Printf("ApiTransactionRequest is : %+v",json)
		log.Printf("paramTransaction is : %s",strings.Join(paramTransaction,","))
		//log.Println(fmt.Sprintf("paramTransaction is : %s",strings.Join(paramTransaction,",")))

		c.JSON(http.StatusOK,apiTransactiondataResponse)
	})

	r.Run()// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}