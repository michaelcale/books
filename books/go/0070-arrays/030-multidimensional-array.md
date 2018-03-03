---
Title: Multidimensional array
Id: 7303
Score: -2
---
Multidimensional arrays are basically arrays containing other arrays as elements.
It is represented like `[sizeDim1][sizeDim2]..[sizeLastDim]type`, replacing `sizeDim` by numbers corresponding to the length of the dimension, and `type` by the type of data in the multidimensional array.

For example, `[2][3]int` represents an array composed of **2 sub arrays** of **3 int typed elements**.
This could be used to represent a matrix of **2 lines** and **3 columns**.

It's possible to create arrays with a lot of dimensions. For example `var values := [2017][12][31][24][60]int` could store a number for each minute since Year 0.

To access the array in the example above (for example, searching for the value of "2016-01-31 at 19:42"), you would access `values[2016][0][30][19][42]` (because **array indexes start at 0** and not at 1 like days and months)

Some examples follow:

@file multidimension_array.go output sha1:6f6d04c233a8797b3f8e844dd90333e9b113a6c7 goplayground:50YPzcT_N8u

@file multidimension_array2.go sha1:36bc4a283726d10519e47ffb3718bbcd3b612de1 goplayground:ZOrLrGfZfW7

@file multidimension_array3.go sha1:8700f6b7766aa277534f63def9cd8a1cbe4e8a21 goplayground:iTvnKtbjN9L
