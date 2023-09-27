package coin

/*
CNY -> 人民币
Coin -> Mi币
*/
func CNY2Coin(cny float64) int64 {
	return int64(cny * 10)
}

func Coin2CNY(coin int64) float64 {
	return float64(coin) / 10
}
