package godd

import (
	"os"
)

type Duplicator struct {
	inFile  *os.File
	outFile *os.File
	bs      int
	count   int
}

func NewDuplicator(bs, count int, input, output string) (*Duplicator, error) {

	//fmt.Printf("")
	//fmt.Printf("blockSize - %d\n", bs)
	//fmt.Printf("count - %d\n", count)
	//fmt.Printf("inputFile - %s\n", input)
	//fmt.Printf("outputFile - %s\n", output)
	//
	//if input != "" {
	//	inputFile, err := os.OpenFile(input, os.O_RDONLY, 0755)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//} else {
	//	inputFile := os.Stdin
	//}
	//defer inputFile.Close()
	//
	//if output != "" {
	//	outputFile, err := os.OpenFile(output, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	//	if err != nil {
	//		return err
	//	}
	//} else {
	//	outputFile := os.Stdout
	//}
	//defer outputFile.Close()
	//
	//
	//fmt.Printf("Works!")
	//return nil
	return nil, nil
}

func (d *Duplicator) setBlockSize(bs int) {
	if bs == 0 {
		d.bs = 1024
	} else {
		d.bs = bs
	}
}

func (d *Duplicator) ReadCount(in, out *os.File, read chan<- int) error {
	return nil
}

func (d *Duplicator) readFile(in, out *os.File, read chan<- int, bs, count int) error {
	return nil
}

func (d *Duplicator) DoCool() {

}
