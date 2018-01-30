Title: Timed loop
Id: 16668
Score: 2
Body:
    package main
    
    import(
        "fmt"
        "time"
    )
    
    func main() {
        for _ = range time.Tick(time.Second * 3) {
            fmt.Println("Ticking every 3 seconds")
        }
    }
|======|
