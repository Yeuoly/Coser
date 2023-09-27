package typ

import (
	"fmt"
	"strconv"
)

/*
Convert interface{} to string
Example:

	InterfaceToString(1) -> "1"
	InterfaceToString("1") -> "1"
*/
func InterfaceToString(i interface{}) string {
	switch typ := i.(type) {
	case string:
		return i.(string)
	case int:
		return strconv.Itoa(i.(int))
	case int16:
		return strconv.Itoa(int(i.(int16)))
	case int32:
		return strconv.Itoa(int(i.(int32)))
	case int64:
		return strconv.Itoa(int(i.(int64)))
	case uint:
		return strconv.Itoa(int(i.(uint)))
	case uint16:
		return strconv.Itoa(int(i.(uint16)))
	case uint32:
		return strconv.Itoa(int(i.(uint32)))
	case uint64:
		return strconv.Itoa(int(i.(uint64)))
	case float32:
		return strconv.FormatFloat(float64(i.(float32)), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64)
	case []byte:
		return string(i.([]byte))
	default:
		return fmt.Sprintf("%v", typ)
	}
}

/*
 * GetPartialType returns the type of the interface{}.
 * mostly, go standard library will return the full type of the interface{}. such as "*types.User"
 * Example:
 * 	type User struct {}
 * 	GetPartialType(User{}) -> "User"
 * 	GetPartialType(&User{}) -> "User"
 *
 */
func GetPartialType(i interface{}) string {
	return fmt.Sprintf("%T", i)
}
