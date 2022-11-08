package main

import (
	"context"
	"log"
)

func main() {


	ctx,cancel:=context.WithCancel(context.Background())

	cancel()

	testfunc(ctx)

	
	log.Println("after testfunc")

	for i := 0; i < 10; i++ {
		log.Println(i,"in main")
	}
	
	//fmt.Scanln()

}

func testfunc(ctx context.Context) {
	log.Println("in testfunc")

	c:=make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c<-i
		}
	}()


	for i := 0; i < 10; i++ {
		log.Println("receiving value",<-c)
	}

	// var temp int 

	// select{
	// 	case<-ctx.Done():
	// 		log.Println("context done")
	// 		return
	// 	case temp=<-c:
	// 		log.Println("receiving value",temp)
	// }



}