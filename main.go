package main

import (
	"flag"
	"fmt"
	"github.com/omerkaya1/godd/cmd/godd"
	"log"
)

func main() {
	// Available commands
	blockSize := flag.Int64("bs", 1024, "read and write up to BYTES bytes at a time (default: 1024)")
	offset := flag.Int64("offset", 0, "read from the specified position in BYTES")
	count := flag.Int("count", 0, "copy only N input blocks")
	inputFile := flag.String("if", "", "read from FILE instead of stdin")
	outputFile := flag.String("of", "", "write to FILE instead of stdout")

	fmt.Printf("%s\n", *inputFile)
	fmt.Printf("%s\n", *outputFile)

	flag.Parse()
	d, err := godd.NewDuplicator(blockSize, offset, count, inputFile, outputFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := d.DoCool(); err != nil {
		log.Fatal(err)
		return
	}
}
