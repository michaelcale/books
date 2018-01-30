Title: Ping pong with two goroutines
Id: 6056
Score: 2
Body:
    package main

    import (
        "fmt"
        "time"
    )

    // The pinger prints a ping and waits for a pong
    func pinger(pinger <-chan int, ponger chan<- int) {
        for {
            <-pinger
            fmt.Println("ping")
            time.Sleep(time.Second)
            ponger <- 1
        }
    }

    // The ponger prints a pong and waits for a ping
    func ponger(pinger chan<- int, ponger <-chan int) {
        for {
            <-ponger
            fmt.Println("pong")
            time.Sleep(time.Second)
            pinger <- 1
        }
    }

    func main() {
        ping := make(chan int)
        pong := make(chan int)

        go pinger(ping, pong)
        go ponger(ping, pong)

        // The main goroutine starts the ping/pong by sending into the ping channel
        ping <- 1

        for {
            // Block the main thread until an interrupt
            time.Sleep(time.Second)
        }
    }

[Run a slightly modified version of this code in Go Playground][1]


  [1]: https://play.golang.org/p/LXcPiIPrgf
|======|
