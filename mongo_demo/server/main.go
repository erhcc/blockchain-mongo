package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	//"log"
	"time"

	"github.com/erichuang-code/blockchain-mongo/models"
	"github.com/erichuang-code/blockchain-mongo/req/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	filename "github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
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

	url:=viper.GetString("mongodb.url")
	//log.Println(url)
	clientOptions:=options.Client().ApplyURI(url)//mongodb://localhost:27017/
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

//var log=log.New()

// type envvars struct{
// 	Enva string
// }

func setupViper()  {
	viper.SetConfigName("config") // name of config file (without extension)
	//viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.SetConfigType("json") 
	viper.AddConfigPath("config/")   // path to look for the config file in
	// viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	url:=viper.GetString("mongodb.url")
	log.Printf("viper read config succ,and url is %s",url)
	// envtest:=viper.GetString("envtest")

	// env:=envvars{}

	// viper.BindEnv("Enva")
	// viper.Unmarshal(&env)

	// log.Printf("os.get envtest, value is: %s",os.Getenv("envtest"))

	// log.Printf("os.get GOPATH, value is: %s,-> using viper, get value:%s",os.Getenv("GOPATH"),viper.GetString("GOPATH"))

	// log.Printf("viper read config succ,and url is %s,envtest is %s,bindenv test %+v",url,envtest,env)
}

func setupLogger()  {
	formatter:=new(log.TextFormatter)
	formatter.FullTimestamp=true
	formatter.ForceColors=true

	// formatter:=new(log.JSONFormatter)


	log.AddHook(filename.NewHook())//print filename+line at every log
	log.SetFormatter(formatter)
	//formatter.TimestampFormat

	  // You could set this to any `io.Writer` such as a file
  file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  if err == nil {
   //log.Out = file
   log.SetOutput(file)
  } else {
   log.Info("Failed to log to file, using default stderr")
  }

}

func main(){

	setupViper()
	setupLogger()

    ticker := time.NewTicker(2000 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                log.Println("Tick at", t)
				handleBlockAndTransactionData()
				log.Println("******done******")
            }
        }
    }()

    time.Sleep(6600 * time.Millisecond)
    ticker.Stop()
    done <- true
    log.Println("Ticker stopped")

	fmt.Scanln()
}

func handleBlockAndTransactionData(){
	log.Printf("query block\n")

	resJsonByteArrayBlock,err:=queryBlock()
	//var user models.User
	if err==nil {
		//json.Unmarshal(model,&user)
		var apiBlockdataResponse models.ApiBlockdataResponse
		json.Unmarshal(resJsonByteArrayBlock,&apiBlockdataResponse)

		log.Printf("apiBlockdataResponse is %s\n",string(resJsonByteArrayBlock))
		//save in mongoDB
		initDb()
		collection.InsertOne(ctx,apiBlockdataResponse)
		log.Printf("block data saved\n")

		//query transactions
		transactions:=apiBlockdataResponse.Result.Transactions

		var paramTransaction=transactions
		var resultJsonByteArrayTransaction []byte
		resultJsonByteArrayTransaction,err=queryTransaction(paramTransaction)

		var apiTransactiondataResponse models.ApiTransactiondataResponse
		json.Unmarshal(resultJsonByteArrayTransaction,&apiTransactiondataResponse)

		log.Printf("apiTransactiondataResponse is %+v",apiTransactiondataResponse)
		collection.InsertOne(ctx,apiTransactiondataResponse)

		log.Printf("transcation data saved\n")
		
		// //query each
		// for _,txid:=range transactions{

		// }


	}else{
		log.Printf("err:%v",err)
	}
}

//here for a test
const urlApiDataForPostTest string="https://reqres.in/api/users"

//mock by gin for real block
const urlApiBlock string="http://localhost:8080/blockdata"
const urlApiTransaction string="http://localhost:8080/transactiondata"

func queryBlock()([]byte,error){
/*
{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params": ["0x5BAD55",false],"id":1}
*/
	contentBody:=models.ApiBlockRequest{
		Jsonrpc: "2.0",
		Method: "eth_getBlockByNumber",
		Params: []interface{}{"0x5BAD55",false},
		Id: 1,
	}

	resultByteArray,err:=http.PostJson(urlApiBlock,contentBody)

	if err!=nil {
		log.Printf("err is: %v",err)
		return nil,err
	}else{
		return resultByteArray,nil
	}

}

func queryTransaction(params []string)([]byte,error){
	/*
	'{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params": ["0xbb3a336e3f823ec18197f1e13ee875700f08f03e2cab75f0d0b118dabb44cba0"],"id":1}'
	*/
		contentBody:=models.ApiTransactionRequest{
			Jsonrpc: "2.0",
			Method: "eth_getTransactionByHash",
			Params: params,
			Id: 1,//todo: later when apionline
		}
	
		//query api
		resultByteArray,err:=http.PostJson(urlApiTransaction,contentBody)
	
		if err!=nil {
			fmt.Println(err)
			return nil,err
		}else{
			return resultByteArray,nil
		}
	
	}

