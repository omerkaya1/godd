package main

import (
	"flag"
	"log"

	"github.com/omerkaya1/godd/cmd/godd"
)

func main() {
	// Available commands
	blockSize := flag.Int64("bs", 1024, "read and write up to BYTES bytes at a time (default: 1024)")
	offset := flag.Int64("offset", 0, "read from the specified position in BYTES")
	count := flag.Int64("count", 0, "copy only N input blocks")
	inputFile := flag.String("if", "", "read from FILE instead of stdin")
	outputFile := flag.String("of", "", "write to FILE instead of stdout")

	flag.Parse()

	d, err := godd.NewDuplicator(*blockSize, *offset, *count, *inputFile, *outputFile)
	if err != nil {
		log.Fatal(err)
	}
	if err := d.CopyContents(); err != nil {
		log.Fatal(err)
	}
}
