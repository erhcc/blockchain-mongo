package utils

import (
	"net/http"
	"time"
)

var ClientSingleton http.Client

func init()  {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 1024
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100//2
		
	// httpClient := &http.Client{
	//   Timeout:   10 * time.Second,
	//   Transport: t,
	// }
	ClientSingleton= http.Client{
		Transport: t,
		Timeout:   10 * time.Second,
	}
}
