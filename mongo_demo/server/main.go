package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/erichuang-code/blockchain-mongo/models"
	"github.com/erichuang-code/blockchain-mongo/req/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* funcs:
1. query api
	notes:
		api now provided by internal, aslo can use geth to get data
2. save into mango
*/

var collection *mongo.Collection
var ctx=context.TODO()

func initDb(){

	clientOptions:=options.Client().ApplyURI("mongodb://localhost:27017/")
	client,err:=mongo.Connect(ctx,clientOptions)
	if err!=nil {
		log.Fatal(err)
	}

	err=client.Ping(ctx,nil)
	if err!=nil {
		log.Fatal(err)
	}

	collection=client.Database("local").Collection("test1")
	//insertMany()
}


func main(){
	model,err:=queryApi()
	//var user models.User
	if err==nil {
		//json.Unmarshal(model,&user)
		var apiBlockdataResponse models.ApiBlockdataResponse
		json.Unmarshal(model,&apiBlockdataResponse)
		fmt.Println("output in main below:")
		fmt.Println(string(model))
		//save in mongoDB
		initDb()
		collection.InsertOne(ctx,apiBlockdataResponse)

	}
}

//here for a test
const urlApiDataForPostTest string="https://reqres.in/api/users"

//mock by gin for real block
const urlApiData string="http://localhost:8080/blockdata"


func queryApi()([]byte,error){
/*
{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params": ["0x5BAD55",false],"id":1}
*/
	contentBody:=http.ReqBlockChain{
		Jsonrpc: "2.0",
		Method: "eth_getBlockByNumber",
		Params: []interface{}{"0x5BAD55",false},
		Id: 1,
	}

	resultByteArray,err:=http.PostJson(urlApiData,contentBody)

	if err!=nil {
		fmt.Println(err)
		return nil,err
	}else{
		return resultByteArray,nil
	}

}

