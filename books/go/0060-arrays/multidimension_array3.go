package main

import "fmt"

func main() {
	// :show start

	var multiDimArray = [2][4][3][2]string{}

	// We can set some values in the array's cells
	multiDimArray[0][0][0][0] = "All zero indexes"   // Setting the first value
	multiDimArray[1][3][2][1] = "All indexes to max" // Setting the value at extreme location

	// The data looks like:
	// > [[[["All zero indexes" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
	//    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
	//    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
	//    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]]
	//   [[[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
	//    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
	//    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
	//    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" "All indexes to max"]]]]

	// :show end

	fmt.Printf("%#v\n", multiDimArray)
}
