Title: Simple Select Working with Channels
Id: 12206
Score: 1
Body:
In this example we create a goroutine (a function running in a separate thread) that accepts a `chan` parameter, and simply loops, sending information into the channel each time.

In the `main` we have a `for` loop and a `select`. The `select` will block processing until one of the `case` statements becomes true. Here we have declared two cases; the first is when information comes through the channel, and the other is if no other case occurs, which is known as `default`.

    // Use of the select statement with channels (no timeouts)
    package main
    
    import (
        "fmt"
        "time"
    )
    
    // Function that is "chatty"
    // Takes a single parameter a channel to send messages down
    func chatter(chatChannel chan<- string) {
        // Clean up our channel when we are done.
        // The channel writer should always be the one to close a channel.
        defer close(chatChannel)
    
        // loop five times and die
        for i := 1; i <= 5; i++ {
            time.Sleep(2 * time.Second) // sleep for 2 seconds
            chatChannel <- fmt.Sprintf("This is pass number %d of chatter", i)
        }
    }
    
    // Our main function
    func main() {
        // Create the channel
        chatChannel := make(chan string, 1)
    
        // start a go routine with chatter (separate, non blocking)
        go chatter(chatChannel)
    
        // This for loop keeps things going while the chatter is sleeping
        for {
            // select statement will block this thread until one of the two conditions below is met
            // because we have a default, we will hit default any time the chatter isn't chatting
            select {
            // anytime the chatter chats, we'll catch it and output it
            case spam, ok := <-chatChannel:
                // Print the string from the channel, unless the channel is closed
                // and we're out of data, in which case exit.
                if ok {
                    fmt.Println(spam)
                } else {
                    fmt.Println("Channel closed, exiting!")
                    return
                }
            default:
                // print a line, then sleep for 1 second.
                fmt.Println("Nothing happened this second.")
                time.Sleep(1 * time.Second)
            }
        }
    }

[Try it on the Go Playground!](https://play.golang.org/p/jMeu32yIUJ)
|======|
