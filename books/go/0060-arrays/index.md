Title: Arrays
Id: 390
Introduction:
Arrays are specific data type, representing an ordered collection of elements of another type.

In Go, Arrays can be simple (sometime called "lists") or multi-dimensional (like for example a 2-dimentions arrays is representing a ordered collection of arrays, that contains elements)
|======|
Syntax:
 - var variableName [5]ArrayType // Declaring an array of size 5.
 - var variableName [2][3]ArrayType = { {Value1, Value2, Value3}, {Value4, Value5, Value6} }  // Declaring a multidimensional array
 - variableName := [...]ArrayType {Value1, Value2, Value3} // Declare an array of size 3 (The compiler will count the array elements to define the size)
 - arrayName[2]              // Getting the value by index.
 - arrayName[5] = 0          // Setting the value at index.
 - arrayName[0]              // First value of the Array
 - arrayName[ len(arrayName)-1 ] // Last value of the Array
|======|
