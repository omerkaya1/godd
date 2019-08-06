package godd

import (
	"fmt"
	"io"
	"os"
)

// Duplicator is a type that
type Duplicator struct {
	inFile  *os.File
	outFile *os.File
	bs      int64
	count   int
	offset  int64
}

// NewDuplicator returns an
func NewDuplicator(bs, offset *int64, count *int, input, output *string) (*Duplicator, error) {
	// Main variables
	var inFile, outFile *os.File
	var err error

	if *input == "" {
		inFile = os.Stdin
	} else {
		inFile, err = os.OpenFile(*input, os.O_RDONLY, 0644)
		if err != nil {
			return nil, err
		}
	}

	if *output == "" {
		outFile = os.Stdout
	} else {
		outFile, err = os.OpenFile(*output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return nil, err
		}
	}

	return &Duplicator{
		inFile:  inFile,
		outFile: outFile,
		bs:      setBlockSize(*bs),
		count:   *count,
		offset:  *offset,
	}, nil
}

func setBlockSize(bs int64) int64 {
	if bs == 0 {
		return 1024
	} else {
		return bs
	}
}

//func (d *Duplicator) ReadCount(in, out *os.File, read chan<- int) error {
//	return nil
//}
//
//func (d *Duplicator) readFile(in, out *os.File, read chan<- int, bs, count int) error {
//	return nil
//}

func (d *Duplicator) DoCool() error {
	defer d.inFile.Close()
	defer d.outFile.Close()
	var total int64
	buf := make([]byte, d.bs)

	// Get the file statistics
	stat, err := d.inFile.Stat()
	if err != nil {
		return err
	}

	// Set the offset (if 0, then it'll start from the beginning of the file)
	if d.offset > 0 {
		if _, err := d.inFile.Seek(d.offset, 0); err != nil {
			return err
		}
	}

	if !stat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", stat.Name())
	}

	if _, err := io.WriteString(os.Stdout, "Copying started!\n"); err != nil {
		return err
	}

	for {
		n, err := d.inFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		_, err = d.outFile.Write(buf[:n])
		if err != nil {
			return err
		}
		total += int64(n)
	}

	if _, err := fmt.Fprintf(os.Stdout, "[%d / %d] copied...\n", total, stat.Size()); err != nil {
		return err
	}

	if _, err := io.WriteString(os.Stdout, "Done!\n"); err != nil {
		return err
	}

	return nil
}
