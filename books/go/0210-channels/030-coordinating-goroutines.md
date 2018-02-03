Title: Coordinating goroutines
Id: 9050
Score: 0
Body:
Imagine a goroutine with a two step process, where the main thread needs to do some work between each step:

```
func main() {
    ch := make(chan struct{})
    go func() {
        // Wait for main thread's signal to begin step one
        <-ch
        
        // Perform work
        time.Sleep(1 * time.Second)
        
        // Signal to main thread that step one has completed
        ch <- struct{}{}
        
        // Wait for main thread's signal to begin step two
        <-ch
        
        // Perform work
        time.Sleep(1 * time.Second)
        
        // Signal to main thread that work has completed
        ch <- struct{}{}
    }()
    
    // Notify goroutine that step one can begin
    ch <- struct{}{}
    
    // Wait for notification from goroutine that step one has completed
    <-ch

    // Perform some work before we notify
    // the goroutine that step two can begin
    time.Sleep(1 * time.Second)
    
    // Notify goroutine that step two can begin
    ch <- struct{}{}
    
    // Wait for notification from goroutine that step two has completed
    <-ch
}
```
|======|
