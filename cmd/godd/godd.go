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
	count   int64
	offset  int64
}

// NewDuplicator returns a new Duplicator object
func NewDuplicator(bs, offset, count int64, input, output string) (*Duplicator, error) {
	// Main variables
	var inFile, outFile *os.File
	var err error

	if input == "" {
		inFile = os.Stdin
	} else {
		inFile, err = os.Open(input)
		if err != nil {
			return nil, err
		}
	}

	if output == "" {
		outFile = os.Stdout
	} else {
		outFile, err = os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return nil, err
		}
	}

	return &Duplicator{
		inFile:  inFile,
		outFile: outFile,
		bs:      bs,
		count:   count,
		offset:  offset,
	}, nil
}

func (d *Duplicator) CopyContents() error {
	defer d.inFile.Close()
	defer d.outFile.Close()
	// Get the input file statistics
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

	// Check whether the input is a regular file or STDIN
	if !stat.Mode().IsRegular() {
		if _, err := io.WriteString(os.Stdout, "Input text to copy below. (to exit press Ctrl+D)\n"); err != nil {
			return err
		}
	} else {
		if _, err := io.WriteString(os.Stdout, "Copying started...\n"); err != nil {
			return err
		}
	}

	// Progress report part
	if _, err := io.WriteString(os.Stdout, fmt.Sprintf("Bytes to copy %d\n", stat.Size())); err != nil {
		return err
	}

	total, err := d.writeToOutput()
	if err != nil {
		return fmt.Errorf("%v\nWritten: %d bytes", err, total)
	}

	if _, err := fmt.Fprintf(os.Stdout, "[%d / %d] copied in total\n", total, stat.Size()); err != nil {
		return err
	}
	return nil
}

func (d *Duplicator) writeToOutput() (int64, error) {
	var total, counter int64
	buf := make([]byte, d.bs)

	// Set the limit to write from the source
	if d.count > 0 {
		counter = d.count
	}

	for counter > 0 {
		n, err := d.inFile.Read(buf)
		if err != nil && err != io.EOF {
			return total, err
		}
		if n == 0 || err == io.EOF {
			break
		}
		_, err = d.outFile.Write(buf[:n])
		if err != nil {
			return total, err
		}
		// We check for count aka limit, so that writing could be stopped
		if int64(n) == d.bs && d.count > 0 {
			counter--
		}
		total += int64(n)
	}

	return total, nil
}
