---
Title: Basic cpu and memory profiling
Id: 259
Score: 1
SOId: 25406
---
Add the following code in you main program.
```go
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }
    ...
    if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
        f.Close()
    }
}
```

After that **build** the go program if added in main `go build main.go`. Run main program with flags defined in code `main.exe -cpuprofile cpu.prof -memprof mem.prof`. If the profiling is done for test cases run the tests with same flags `go test -cpuprofile cpu.prof -memprofile mem.prof`
