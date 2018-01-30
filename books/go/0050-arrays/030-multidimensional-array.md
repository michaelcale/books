Title: Multidimensional Array
Id: 7303
Score: -2
Body:
Multidimensional arrays are basically arrays containing others arrays as elements.  
It is represented like `[sizeDim1][sizeDim2]..[sizeLastDim]type`, replacing `sizeDim` by numbers corresponding to the length of the dimention, and `type` by the type of data in the multidimensional array.

For example, `[2][3]int` is representing an array composed of **2 sub arrays** of **3 int typed elements**.  
It can basically be the representation of a matrix of **2 lines** and **3 columns**. 

So we can make huge dimensions number array like `var values := [2017][12][31][24][60]int` for example if you need to store a number for each minutes since Year 0.

To access this kind of array, for the last example, searching for the value of 2016-01-31 at 19:42, you will access `values[2016][0][30][19][42]` (because **array indexes starts at 0** and not at 1 like days and months)

Some examples following:

    // Defining a 2d Array to represent a matrix like
    // 1 2 3     So with 2 lines and 3 columns;
    // 4 5 6     
    var multiDimArray := [2/*lines*/][3/*columns*/]int{ [3]int{1, 2, 3}, [3]int{4, 5, 6} }

    // That can be simplified like this:
    var simplified := [2][3]int{{1, 2, 3}, {4, 5, 6}}

    // What does it looks like ?
    fmt.Println(multiDimArray)
    // > [[1 2 3] [4 5 6]]

    fmt.Println(multiDimArray[0]) 
    // > [1 2 3]    (first line of the array)

    fmt.Println(multiDimArray[0][1])
    // > 2          (cell of line 0 (the first one), column 1 (the 2nd one))

<!-- break -->

    // We can also define array with as much dimensions as we need
    // here, initialized with all zeros
    var multiDimArray := [2][4][3][2]string{} 
  
    fmt.Println(multiDimArray);
    // Yeah, many dimensions stores many data
    // > [[[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
    //    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
    //    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
    //    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]]
    //   [[[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
    //    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
    //    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
    //    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]]

<!-- break -->
    // We can set some values in the array's cells
    multiDimArray[0][0][0][0] := "All zero indexes"  // Setting the first value
    multiDimArray[1][3][2][1] := "All indexes to max"  // Setting the value at extreme location

    fmt.Println(multiDimArray);
    // If we could see in 4 dimensions, maybe we could see the result as a simple format
        
    // > [[[["All zero indexes" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
    //    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
    //    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
    //    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]]
    //   [[[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
    //    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
    //    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" ""]]]
    //    [[["" ""] ["" ""]] [["" ""] ["" ""]] [["" ""] ["" "All indexes to max"]]]]
    

|======|
