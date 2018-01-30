Title: OS Signals
Id: 4497
Syntax:
- func Notify(c chan<- os.Signal, sig ...os.Signal)
|======|
Parameters:
|Parameter | Details | 
|----------|--------| 
| c chan<- os.Signal | Receiving `channel` specifically of type `os.Signal`; easily created with `sigChan := make(chan os.Signal)` | 
| sig ...os.Signal | List of `os.Signal` types to catch and send down this `channel`. See https://golang.org/pkg/syscall/#pkg-constants for more options.
|======|
