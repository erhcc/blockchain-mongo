package main

import (
	"net/http"

	"github.com/erichuang-code/blockchain-mongo/models"
	"github.com/gin-gonic/gin"
)

// type blockChainResponse struct{
// 	Result models.Transaction
// }

var data =models.ApiBlockdataResponse{
	Id:      1,
	Jsonrpc: "2.0",
	Result:  models.Result{
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

/* mock:
1. serve api to blockchain
*/
func main(){

	r:=gin.Default()

	r.POST("/blockdata",func (c *gin.Context)  {
		c.JSON(http.StatusOK,data)
	})

	r.Run()// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}