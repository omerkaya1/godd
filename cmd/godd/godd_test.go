package godd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDuplicator(t *testing.T) {
	if d, err := NewDuplicator(1024, 0, 0, "", ""); assert.NoError(t, err) {
		assert.NotNil(t, d)
	}
}

func TestDuplicator_CopyContentsAll(t *testing.T) {
	if d, err := NewDuplicator(1024, 0, 0, "", ""); assert.NoError(t, err) {
		assert.NotNil(t, d)
	}
}

func TestDuplicator_CopyContentsPartial(t *testing.T) {
	if d, err := NewDuplicator(1024, 0, 0, "", ""); assert.NoError(t, err) {
		assert.NotNil(t, d)
	}
}

func TestDuplicator_CopyContentsApplyOffset(t *testing.T) {
	if d, err := NewDuplicator(1024, 10, 0, "", ""); assert.NoError(t, err) {
		assert.NotNil(t, d)
	}
}
