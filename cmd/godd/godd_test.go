package godd

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	offset    int64 = 10
	bs        int64 = 10
	count     int64 = 10
	inFile          = "../../testdata/test_file.txt"
	outFile         = "../../testdata/output_file.txt"
	testInput       = []byte("This is a test string.")
)

func TestNewDuplicator(t *testing.T) {
	if d, err := NewDuplicator(1024, 0, 0, inFile, outFile); assert.NoError(t, err) {
		assert.NotNil(t, d)
	}
}

func TestDuplicator_CopyContentsAll(t *testing.T) {
	if d, err := NewDuplicator(1024, 0, 0, inFile, outFile); assert.NoError(t, err) {
		assert.NotNil(t, d)
		assert.NoError(t, d.CopyContents())
		assert.FileExists(t, outFile)
		if stat1, stat2, err := retStats(true, true); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, stat1.Size(), stat2.Size())
		}
	}
}

func TestDuplicator_CopyContentsPartial(t *testing.T) {
	if d, err := NewDuplicator(bs, 0, count, inFile, outFile); assert.NoError(t, err) {
		assert.NotNil(t, d)
		assert.NoError(t, d.CopyContents())
		assert.FileExists(t, outFile)
		if _, stat2, err := retStats(false, true); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, stat2.Size(), bs*count)
		}
	}
}

func TestDuplicator_CopyContentsApplyOffset(t *testing.T) {
	if d, err := NewDuplicator(bs, offset, 0, inFile, outFile); assert.NoError(t, err) {
		assert.NotNil(t, d)
		assert.NoError(t, d.CopyContents())
		assert.FileExists(t, outFile)
		if stat1, stat2, err := retStats(true, true); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, stat2.Size(), stat1.Size()-offset)
		}
	}
}

func TestDuplicator_CopyContentsStdinToStdout(t *testing.T) {
	if d, err := NewDuplicator(1024, 0, 0, "", ""); assert.NoError(t, err) {
		assert.NotNil(t, d)
		r, w := io.Pipe()
		defer r.Close()
		defer w.Close()

		go func() {
			_, err := w.Write(testInput)
			if err != nil {
				t.Error(err)
			}
		}()

		buf := make([]byte, len(testInput))
		if _, err := r.Read(buf); err != nil {
			t.Error(err)
		}
		assert.Equal(t, testInput, buf)
	}
}

func retStats(in, out bool) (os.FileInfo, os.FileInfo, error) {
	var input, output *os.File
	var stat1, stat2 os.FileInfo
	var err error

	if in {
		input, err = os.Open(inFile)
		if err != nil {
			return nil, nil, err
		}
		defer input.Close()

		stat1, err = input.Stat()
		if err != nil {
			return nil, nil, err
		}
	}

	if out {
		output, err = os.Open(outFile)
		if err != nil {
			return nil, nil, err
		}
		defer output.Close()

		stat2, err = output.Stat()
		if err != nil {
			return nil, nil, err
		}
	}

	return stat1, stat2, nil
}
