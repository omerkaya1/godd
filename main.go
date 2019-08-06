package main

import (
	"flag"
	"github.com/omerkaya1/godd/cmd/godd"
	"log"
)

func main() {
	// Available commands
	blockSize := *flag.Int("bs", 1024, "read and write up to BYTES bytes at a time (default: 1024)")
	count := *flag.Int("count", 0, "copy only N input blocks")
	inputFile := *flag.String("if", "", "read from FILE instead of stdin")
	outputFile := *flag.String("of", "", "write to FILE instead of stdout")

	flag.Parse()
	d, err := godd.NewDuplicator(blockSize, count, inputFile, outputFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	d.DoCool()
	return
}
