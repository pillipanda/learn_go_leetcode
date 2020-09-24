package decode_ways

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeWays(t *testing.T) {
	resp := DecodeWays("611")
	fmt.Println(resp)
	assert.Equal(t, resp, 2)
}
