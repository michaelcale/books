Title: Inline Expansion
Id: 2718
Remarks:
Inline expansion is a common optimization in compiled code that prioritized performance over binary size. It lets the compiler replace a function call with the actual body of the function; effectively copy/pasting code from one place to another at compile time. Since the call site is expanded to just contain the machine instructions that the compiler generated for the function, we don't have to perform a CALL or PUSH (the x86 equivalant of a GOTO statement or a stack frame push) or their equivalant on other architectures.

The inliner makes decisions about whether or not to inline a function based on a number of heuristics, but in general Go inlines by default. Because the inliner gets rid of function calls, it effectively gets to decide where the scheduler is allowed to preempt a goroutine.

Function calls will not be inlined if any of the following are true (there are many other reasons too, this list is incomplete):

  - Functions are variadic (eg. they have `...` args)
  - Functions have a "max hairyness" greater than the budget (they recurse too much or can't be analyzed for some other reason)
  - They contain `panic`, `recover`, or `defer`
|======|
