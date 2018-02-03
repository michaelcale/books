Title: Error Handling
Id: 785
Introduction:
In Go, unexpected situations are handled using **errors**, not exceptions. This approach is more similar to that of C, using errno, than to that of Java or other object-oriented languages, with their try/catch blocks. However, an error is not an integer but an interface.

A function that may fail typically returns an **error** as its last return value. If this error is not **nil**, something went wrong, and the caller of the function should take action accordingly.
|======|
Remarks:
Note how in Go you don't _raise_ an error. Instead, you _return_ an error in case of failure.

If a function can fail, the last returned value is generally an `error` type.

    // This method doesn't fail
    func DoSomethingSafe() {
    }
    
    // This method can fail
    func DoSomething() (error) {
    }
    
    // This method can fail and, when it succeeds,
    // it returns a string.
    func DoAndReturnSomething() (string, error) {
    }


|======|
