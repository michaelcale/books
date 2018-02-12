package main

import "fmt"

func main() {
	// :show start
	// Defining a 2d Array to represent a matrix like
	// 1 2 3     So with 2 lines and 3 columns;
	// 4 5 6
	multiDimArray := [2] /*lines*/ [3] /*columns*/ int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}

	// That can be simplified like this:
	var simplified = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// What does it looks like ?
	fmt.Println(multiDimArray)
	// > [[1 2 3] [4 5 6]]

	fmt.Println(multiDimArray[0])
	// > [1 2 3]    (first line of the array)

	fmt.Println(multiDimArray[0][1])
	// > 2          (cell of line 0 (the first one), column 1 (the 2nd one))

	// :show end

	// silence compiler error about unused variable
	_ = simplified
}
