package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params": ["0x5BAD55",false],"id":1}
*/
type ReqBlockChain struct{
	Jsonrpc string  	`json:"jsonrpc"`
	Method string  	`json:"method"`
	Params []interface{}	`json:"params"`
	Id int  		`json:"id"`
}

func PostJson(url string, contentBody interface{}) (string,error){

	body,_:=json.Marshal(contentBody) 

	resp, err := http.Post(url,"application/json",bytes.NewBuffer(body))

	// An error is returned if something goes wrong
	if err != nil {
		panic(err)
	}
	//Need to close the response stream, once response is read.
	//Hence defer close. It will automatically take care of it.
	defer resp.Body.Close()

	//Check response code, if New user is created then read response.
	if resp.StatusCode == http.StatusCreated {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//Failed to read response.
			//panic(err)
			return "",err
		}

		//Convert bytes to String and print
		jsonStr := string(body)
		//fmt.Println("Response: ", jsonStr)
		return jsonStr,nil

	} else {
		//The status is not Created. print the error.
		//errMsg :="status not correct"+resp.Status
		errMsg :=fmt.Sprintf("Get failed with error: %s",resp.Status)
		//fmt.Println(errMsg)
		return "",errors.New(errMsg)
		
	}

}
