Title: Logging to file
Id: 13467
Score: 0
Body:
It is possible to specify log destination with something that statisfies io.Writer interface. With that we can log to file:

    package main
    
    import (
        "log"
        "os"
    )
    
    func main() {
        logfile, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
        if err != nil {
            log.Fatalln(err)
        }
        defer logfile.Close()
    
        log.SetOutput(logfile)
        log.Println("Log entry")
    }

Output:

    
    $ cat test.log
    2016/07/26 07:29:05 Log entry
|======|
