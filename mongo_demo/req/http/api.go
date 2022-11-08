package http

import (
	"bytes"
	"encoding/json"
	//"time"

	//"time"

	//"errors"
	//"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	//"github.com/erichuang-code/blockchain-mongo/utils/http"
)

func PostJson(url string, contentBody interface{}) ([]byte,error){


	// defer func ()([]byte,error) {
	// 	if r:=recover();r!=nil{
	// 		log.Warnf("recovered, details is %v",r)
	// 		return nil,nil	
	// 	}

	// }()

	defer func () {
		if r:=recover();r!=nil{
			log.Warnf("recovered, details is %v",r)
		}
	}()

	//fmt.s
	//_ :=fmt.Sprintf("Get failed with error: %s",url)
	
	body,_:=json.Marshal(contentBody) 

	//resp, err := http.Post(url,"application/json",bytes.NewBuffer(body))

	//resp, err := utils.DoHttpReq(url,time.Second*10,nil,bytes.NewBuffer(body))

	//this is used for thirdparty. it could be working for 1000 
	//httpClient:=utils.GetClient(time.Second*60)//utils.GetClient(time.Second*10) utils.ClientSingleton
	resp, err := http.Post(url,"application/json",bytes.NewBuffer(body))

	// An error is returned if something goes wrong
	if err != nil {
		log.Printf("contentbody:%+v,err is :%v",string(body),err)
		//panic(err)
		//resp.Body.Close()
		return nil,err
	}
	//Need to close the response stream, once response is read.
	//Hence defer close. It will automatically take care of it.
	defer resp.Body.Close()

	//Check response code, if New user is created then read response.
	//StatusCreated or ok
	//if resp.StatusCode == http.StatusOK {
		resBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//Failed to read response.
			//panic(err)
			return nil,err
		}

		//Convert bytes to String and print
		//jsonStr := string(body)
		//fmt.Println("Response: ", jsonStr)
		return resBody,nil

	// } else {
	// 	//The status is not Created. print the error.
	// 	//errMsg :="status not correct"+resp.Status
	// 	errMsg :=fmt.Sprintf("Get failed with error: %s",resp.Status)
	// 	//fmt.Println(errMsg)
	// 	return nil,errors.New(errMsg)
		
	// }

}
