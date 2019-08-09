package godd

import (
	"fmt"
	"io"
	"os"
)

// Duplicator is a type that holds all the data.
type Duplicator struct {
	inFile  string
	outFile string
	bs      int64
	count   int64
	offset  int64
}

// NewDuplicator returns a new Duplicator object.
func NewDuplicator(bs, offset, count int64, input, output string) (*Duplicator, error) {
	return &Duplicator{
		inFile:  input,
		outFile: output,
		bs:      bs,
		count:   count,
		offset:  offset,
	}, nil
}

// CopyContents method, surprisingly, copies contents of the input file to the output file.
func (d *Duplicator) CopyContents() error {
	// Get files
	inFile, outFile, err := d.initFiles()
	if err != nil {
		return err
	}
	defer inFile.Close()
	defer outFile.Close()
	// Get the input file statistics
	stat, err := inFile.Stat()
	if err != nil {
		return err
	}

	// Set the offset (if 0, then it'll start from the beginning of the file).
	if d.offset > 0 {
		if _, err := inFile.Seek(d.offset, 0); err != nil {
			return err
		}
	}

	// Check whether the input is a regular file or STDIN.
	if !stat.Mode().IsRegular() {
		if _, err := io.WriteString(os.Stdout, "Input text to copy below. (to exit press Ctrl+D)\n"); err != nil {
			return err
		}
	} else {
		if _, err := io.WriteString(os.Stdout, "Copying started...\n"); err != nil {
			return err
		}
	}

	// Progress report part.
	if _, err := io.WriteString(os.Stdout, fmt.Sprintf("Bytes to copy %d\n", stat.Size())); err != nil {
		return err
	}

	total, err := d.writeToOutput(inFile, outFile)
	if err != nil {
		return fmt.Errorf("%v\nWritten: %d bytes", err, total)
	}

	if _, err := fmt.Fprintf(os.Stdout, "[%d / %d] copied in total\n", total, stat.Size()); err != nil {
		return err
	}
	return nil
}

func (d *Duplicator) initFiles() (*os.File, *os.File, error) {
	// Main variables
	var inFile, outFile *os.File
	var err error

	if d.inFile == "" {
		inFile = os.Stdin
	} else {
		inFile, err = os.Open(d.inFile)
		if err != nil {
			return nil, nil, err
		}
	}

	if d.outFile == "" {
		outFile = os.Stdout
	} else {
		outFile, err = os.OpenFile(d.outFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return nil, nil, err
		}
	}
	return inFile, outFile, nil
}

func (d *Duplicator) writeToOutput(inFile, outFile *os.File) (int64, error) {
	var total, counter int64
	buf := make([]byte, d.bs)

	// Set the limit to write from the source
	if d.count > 0 {
		counter = d.count - 1
	}
	for counter >= 0 {
		n, err := inFile.Read(buf)
		if err != nil && err != io.EOF {
			return total, err
		}
		if n == 0 || err == io.EOF {
			break
		}
		_, err = outFile.Write(buf[:n])
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
