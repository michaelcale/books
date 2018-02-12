---
Title: Multidimensional array
Id: 7303
Score: -2
---
Multidimensional arrays are basically arrays containing others arrays as elements.
It is represented like `[sizeDim1][sizeDim2]..[sizeLastDim]type`, replacing `sizeDim` by numbers corresponding to the length of the dimention, and `type` by the type of data in the multidimensional array.

For example, `[2][3]int` is representing an array composed of **2 sub arrays** of **3 int typed elements**.
It can basically be the representation of a matrix of **2 lines** and **3 columns**.

So we can make huge dimensions number array like `var values := [2017][12][31][24][60]int` for example if you need to store a number for each minutes since Year 0.

To access this kind of array, for the last example, searching for the value of 2016-01-31 at 19:42, you will access `values[2016][0][30][19][42]` (because **array indexes starts at 0** and not at 1 like days and months)

Some examples following:

@file multidimension_array.go output

@file multidimension_array2.go

@file multidimension_array3.go
