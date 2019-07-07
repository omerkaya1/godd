package dd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoCool(t *testing.T) {
	assert.NoError(t, DoCool())
}
