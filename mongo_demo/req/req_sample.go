package req

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//If the struct variable names does not match with json attributes 
//then you can define the json attributes actual name after json:attname as shown below. 
type User struct {
	Name string  	`json:"name"`
	Job string 	    `json:"job"`
}

func main(){

	//Create user struct which need to post.
	user := User{
		Name: "Test User",
		Job: "Go lang Developer",
	}

	//Convert User to byte using Json.Marshal
	//Ignoring error. 
	body, _ := json.Marshal(user)

	//Pass new buffer for request with URL to post.
  //This will make a post request and will share the JSON data
	resp, err := http.Post("https://reqres.in/api/users", "application/json", bytes.NewBuffer(body) )

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
			panic(err)
		}

		//Convert bytes to String and print
		jsonStr := string(body)
		fmt.Println("Response: ", jsonStr)

	} else {
		//The status is not Created. print the error.
		fmt.Println("Get failed with error: ", resp.Status)
	}
}