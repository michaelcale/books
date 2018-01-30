Title: Methods
Id: 3890
Syntax:
- func (t T) exampleOne(i int) (n int) { return i } // this function will receive copy of struct
- func (t *T) exampleTwo(i int) (n int) { return i }
// this method will receive pointer to struct and will be able to modify it
|======|
