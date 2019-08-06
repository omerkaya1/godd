package godd

import (
	"bufio"
	"fmt"
	"os"
)

// NOTE: should be run as a goroutine
func showProgress(in <-chan int, total uint64) error {
	var overall uint64
	w := bufio.NewWriter(os.Stdout)
	for read := range in {
		overall += uint64(read)
		if _, err := w.WriteString(fmt.Sprintf("%d out of %d finished", overall, total)); err != nil {
			return err
		}
	}
	return nil
}
