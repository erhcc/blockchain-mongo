package main

import (
	"bytes"
	"log"
)

func main() {

	sliceA:=make([]int,3,10)//[0 0 0]
	sliceB:=make([]int,5,10)//[0 0 0 0 0]

	sliceA=sliceA[:6]
	sliceA[4]=9
	//sliceA={1,2,3,4}

	log.Println(sliceA,sliceB)

	////testBytesBuffer()
	//testBytesBuffer2()
	/* test copy function */
	source := []byte{1, 2, 3}
	dsc := make([]byte, 10)
	dsc = append(dsc, 11)
	log.Println(dsc)

	dsc=[]byte{11,12,13,14}

	n := copy(dsc, source)

	log.Println(n,dsc)
	
}

func testBytesBuffer2(){
	//buf:=[]byte{'a','b','c','d','e','f','g','h','i','j'}

	buf:=make([]byte,10)
	buf=buf[:3]

	buffer:=bytes.NewBuffer(buf)
	log.Println(buffer)

	buffer.Grow(8)

	log.Println(buffer)

}

func testBytesBuffer(){
	buffer:=bytes.NewBuffer(make([]byte, 10))
	log.Println(buffer)

	buffer.Grow(8)

	log.Println(buffer)

}



/*
// tryGrowByReslice is a inlineable version of grow for the fast-case where the
// internal buffer only needs to be resliced.
// It returns the index where bytes should be written and whether it succeeded.
func (b *Buffer) tryGrowByReslice(n int) (int, bool) {
	if l := len(b.buf); n <= cap(b.buf)-l {
		b.buf = b.buf[:l+n]
		return l, true
	}
	return 0, false
}
*/