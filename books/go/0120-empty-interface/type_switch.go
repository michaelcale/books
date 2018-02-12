package main

import (
	"fmt"
	"strconv"
)

// :show start
func smartConvertToInt(iv interface{}) (int, error) {
	// inside case statements, v is of type matching case type
	switch v := iv.(type) {
	case int:
		return v, nil
	case string:
		return strconv.Atoi(v)
	case float64:
		return int(v), nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", iv)
	}
}

func printSmartConvertToInt(iv interface{}) {
	i, err := smartConvertToInt(iv)
	if err != nil {
		fmt.Printf("Failed to convert %#v to int\n", iv)
		return
	}
	fmt.Printf("%#v of type %T converted to %d\n", iv, iv, i)
}

func main() {
	printSmartConvertToInt("5")
	printSmartConvertToInt(4)
	printSmartConvertToInt(int32(8))
	printSmartConvertToInt("not valid int")
}

// :show end
