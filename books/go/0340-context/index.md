Title: Context
Id: 2743
Syntax:
- type CancelFunc func()
- func Background() Context
- func TODO() Context
- func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
- func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
- func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
- func WithValue(parent Context, key interface{}, val interface{})
|======|
Remarks:
The `context` package (in Go 1.7) or the `golang.org/x/net/context` package (Pre 1.7) is an interface for creating contexts that can be used to carry request scoped values and deadlines across API boundaries and between services, as well as a simple implementation of said interface.

aside: the word "context" is loosely used to refer to the entire tree, or to individual leaves in the tree, eg. the actual `context.Context` values.

At a high level, a context is a tree. New leaves are added to the tree when they are constructed (a `context.Context` with a parent value), and leaves are never removed from the tree. Any context has access to all of the values above it (data access only flows upwards), and if any context is canceled its children are also canceled (cancelation signals propogate downwards). The cancel signal is implemented by means of a function that returns a channel which will be closed (readable) when the context is canceled; this makes contexts a very efficient way to implement the [pipeline and cancellation concurrency pattern](https://blog.golang.org/pipelines), or timeouts.

By convention, functions that take a context have the first argument `ctx context.Context`. While this is just a convention, it's one that should be followed since many static analysis tools specifically look for this argument. Since Context is an interface, it's also possible to turn existing context-like data (values that are passed around throughout a request call chain) into a normal Go context and use them in a backwards compatible way just by implementing a few methods. Furthermore, contexts are safe for concurrent access so you can use them from many goroutines (whether they're running on parallel threads or as concurrent coroutines) without fear.

## Further Reading ##

 - https://blog.golang.org/context
|======|
