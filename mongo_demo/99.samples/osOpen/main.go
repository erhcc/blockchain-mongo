package main

import (
	"fmt"
	"io"
	"log"
	"os"
)
/* vs ioutil.read
	os.open return fd, which more flexible
	ioutil.read invoke os.read
*/
func main() {
    f, err := os.Open("file.txt")
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }
    defer f.Close()
    buf := make([]byte, 1024)
    for {
        n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
    }

	fmt.Println("\xF0\x9F\x92\xBF")
     fmt.Println("\xF0\x9F\x8E\xB2")
     fmt.Println("\xF0\x9F\x90\xA8")
     fmt.Println("\xF0\x9F\x90\xA7")
     fmt.Println("\xF0\x9F\x90\xAB")
     fmt.Println("\xF0\x9F\x90\xAC")
}