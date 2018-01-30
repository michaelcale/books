Title: Time
Id: 8860
Introduction:
The Go [`time`](https://golang.org/pkg/time/) package provides functionality for measuring and displaying time.

This package provide a structure `time.Time`, allowing to store and do computations on dates and time.
|======|
Syntax:
- time.Date(2016, time.December, 31, 23, 59, 59, 999, time.UTC) // initialize
- date1 == date2 // returns `true` when the 2 are the same moment
- date1 != date2 // returns `true` when the 2 are different moment
- date1.Before(date2) // returns `true` when the first is strictly before the second
- date1.After(date2) // returns `true` when the first is strictly after the second
|======|
