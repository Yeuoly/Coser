package coordinate

func LL2XY(longitude float64, latitude float64) (x float64, y float64) {
	x = longitude
	y = latitude
	return
}

func XY2LL(x float64, y float64) (longitude float64, latitude float64) {
	longitude = x
	latitude = y
	return
}
