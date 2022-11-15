package main

type readOp int8

// Don't use iota for these, as the values need to correspond with the
// names and comments, which is easier to see when being explicit.
const (
	//opead comments
	opRead      readOp = -1 // Any other read operation.
	opInvalid   readOp = 0  // Non-read operation.
	opReadRune1 readOp = 1  // Read rune of size 1.
	opReadRune2 readOp = 2  // Read rune of size 2.
	opReadRune3 readOp = 3  // Read rune of size 3.
	opReadRune4 readOp = 4  // Read rune of size 4.
)



type writeOp int8


const (
	//opwrite comments
	opWrite      readOp = iota
	opWriteRune1  // Read rune of size 1.
	opWriteRune2  // Read rune of size 2.
	opWriteRune3  // Read rune of size 3.
	opWriteRune4  // Read rune of size 4.
)