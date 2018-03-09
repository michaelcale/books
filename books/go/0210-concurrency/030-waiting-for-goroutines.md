---
Title: Wait for goroutines to finish
Id: 139
SOId: 2490
---

[Go programs end when the `main` function ends](http://golang.org/ref/spec#Program_execution), therefore it is common practice to wait for all goroutines to finish.

A common solution for this is to use a [sync.WaitGroup]( http://golang.org/pkg/sync/#WaitGroup) object.

@file waiting_for_goroutines.go output sha1:232512714a4b2adcb99a7edd243ab6acab9bf040 goplayground:8DqdlC1x5fs

WaitGroup usage in order of execution:

 1. Declaration of global variable. Making it global is the easiest way to make it visible to all functions and methods.
 2. Increasing the counter. This must be done in the main goroutine because there is no guarantee that a newly started goroutine will execute before 4 due to memory model [guarantees](http://golang.org/ref/mem#tmp_5).
 3. Decreasing the counter. This must be done at the exit of a goroutine. By using a deferred call, we make sure that it [will be called whenever function ends](http://golang.org/ref/spec#Defer_statements), no matter how it ends.
 4. Waiting for the counter to reach 0. This must be done in the main goroutine to prevent the program from exiting before all goroutines have finished.

Parameters are evaluated [before starting a new goroutine](http://golang.org/ref/spec#Go_statements).

Thus it is necessary to define their values explicitly before `wg.Add(10)` so that possibly-panicking code will not increase the counter. Adding 10 items to the WaitGroup, so it will wait for 10 items before `wg.Wait` returns the control back to `main()` goroutine. Here, the value of i is defined in the for loop.
