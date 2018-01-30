Title: Simple Command Execution
Id: 3522
Score: 1
Body:
    // Execute a command a capture standard out. exec.Command creates the command
    // and then the chained Output method gets standard out. Use CombinedOutput() 
    // if you want both standard out and standerr output
    out, err := exec.Command("echo", "foo").Output()
    if err != nil {
        log.Fatal(err)
    }
|======|
