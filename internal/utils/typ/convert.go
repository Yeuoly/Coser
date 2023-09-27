package typ

import "unsafe"

func StringTobytes(s string) []byte {
	// Convert the string header to a slice header without copying the data.
	stringHeader := (*stringStruct)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&sliceStruct{
		array: stringHeader.str,
		len:   stringHeader.len,
		cap:   stringHeader.len,
	}))
}

func BytesToString(b []byte) string {
	// Convert the slice header to a string header without copying the data.
	sliceHeader := (*sliceStruct)(unsafe.Pointer(&b))
	return *(*string)(unsafe.Pointer(&stringStruct{
		str: sliceHeader.array,
		len: sliceHeader.len,
	}))
}

// stringStruct is the internal representation of a string in Go.
type stringStruct struct {
	str unsafe.Pointer
	len int
}

// sliceStruct is the internal representation of a slice in Go.
type sliceStruct struct {
	array unsafe.Pointer
	len   int
	cap   int
}
