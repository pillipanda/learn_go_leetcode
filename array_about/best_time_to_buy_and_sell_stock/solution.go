package best_time_to_buy_and_sell_stock

import "math"

func BestTimeToBuyAndSell(priceSlice []int) int {
	var minPrice int = math.MaxInt8
	var maxProfile int = 0

	for _, i := range priceSlice {
		if minPrice > i {
			minPrice = i
		}
		if maxProfile < (i - minPrice) {
			maxProfile = i - minPrice
		}
	}

	return maxProfile
}