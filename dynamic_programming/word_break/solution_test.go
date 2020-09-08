package word_break

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	resp := WordBreak("catsandogcat", []string{"cats", "dog", "sand", "and", "cat", "an"})
	fmt.Println(resp)
	assert.Equal(t, resp, true)
}
