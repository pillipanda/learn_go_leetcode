package best_time_to_buy_and_sell_stock

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBestTimeToBuyAndSell(t *testing.T)  {
	priceSlice := []int{7,1,5,3,6,4}
	target := 5

	resp := BestTimeToBuyAndSell(priceSlice)
	assert.Equal(t, resp, target)

	priceSlice = []int{7,6,4,3,1}
	target = 0
	resp = BestTimeToBuyAndSell(priceSlice)
	assert.Equal(t, resp, target)
}