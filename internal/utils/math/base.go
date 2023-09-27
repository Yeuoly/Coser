package math

type Integer interface {
	uint | int | uint8 | int8 | uint16 | int16 | uint32 | int32 | uint64 | int64
}

type Float interface {
	float32 | float64
}

func Abs[T Integer | Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
