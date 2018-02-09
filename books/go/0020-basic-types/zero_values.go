package main

import "fmt"

func main() {
	// :show start
	var zeroBool bool             // = false
	var zeroInt int               // = 0
	var zeroF32 float32           // = 0.0
	var zeroStr string            // = ""
	var zeroPtr *int              // = nil
	var zeroSlice []uint32        // = nil
	var zeroMap map[string]int    // = nil
	var zeroInterface interface{} // = nil
	var zeroChan chan bool        // = nil
	var zeroArray [5]int          // = [0, 0, 0, 0, 0]
	type struc struct {
		a int
		b string
	}
	var zeroStruct struc    // = { a: 0, b: ""}
	var zeroFunc func(bool) // = nil

	fmt.Printf("zeroBool: %v\n", zeroBool)
	fmt.Printf("zeroInt: %v\n", zeroInt)
	fmt.Printf("zeroF32: %v\n", zeroF32)
	fmt.Printf("zeroStr: %#v\n", zeroStr)
	fmt.Printf("zeroPtr: %v\n", zeroPtr)
	fmt.Printf("zeroSlice: %#v\n", zeroSlice)
	fmt.Printf("zeroMap: %#v\n", zeroMap)
	fmt.Printf("zeroInterface: %v\n", zeroInterface)
	fmt.Printf("zeroChan: %v\n", zeroChan)
	fmt.Printf("zeroArray: %v\n", zeroArray)
	fmt.Printf("zeroStruct: %#v\n", zeroStruct)
	fmt.Printf("zeroFunc: %v\n", zeroFunc)
	// :show end
}
