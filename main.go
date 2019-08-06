package main

import (
	"flag"
	"fmt"
)

func main() {

	blockSize := flag.Int("bs", 1024, "read and write up to BYTES bytes at a time (default: 1024)")
	count := flag.Int("count", 1024, "copy only N input blocks")
	inputFile := flag.String("if", "", "read from FILE instead of stdin")
	outputFile := flag.String("of", "", "write to FILE instead of stdout")

	flag.Parse()

	fmt.Printf("")
	fmt.Printf("blockSize - %d\n", *blockSize)
	fmt.Printf("count - %d\n", *count)
	fmt.Printf("inputFile - %s\n", *inputFile)
	fmt.Printf("outputFile - %s\n", *outputFile)

	return
}
