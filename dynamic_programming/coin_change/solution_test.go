package coin_change

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestCoinChange(t *testing.T) {
	coins := []int{2}
	money := 3
	target := -1

	resp := CoinChange(coins, money)
	assert.Equal(t, resp, target)
}
