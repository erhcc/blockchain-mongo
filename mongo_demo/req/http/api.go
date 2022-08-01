package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	log "github.com/sirupsen/logrus"

)

func PostJson(url string, contentBody interface{}) ([]byte,error){

	//fmt.s
	//_ :=fmt.Sprintf("Get failed with error: %s",url)
	
	body,_:=json.Marshal(contentBody) 

	resp, err := http.Post(url,"application/json",bytes.NewBuffer(body))

	// An error is returned if something goes wrong
	if err != nil {
		log.Printf("err is :%v",err)
		panic(err)
	}
	//Need to close the response stream, once response is read.
	//Hence defer close. It will automatically take care of it.
	defer resp.Body.Close()

	//Check response code, if New user is created then read response.
	//StatusCreated or ok
	if resp.StatusCode == http.StatusOK {
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

	} else {
		//The status is not Created. print the error.
		//errMsg :="status not correct"+resp.Status
		errMsg :=fmt.Sprintf("Get failed with error: %s",resp.Status)
		//fmt.Println(errMsg)
		return nil,errors.New(errMsg)
		
	}

}
