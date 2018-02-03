Title: Executing a Command then Continue and Wait
Id: 3523
Score: 2
Body:
    cmd := exec.Command("sleep", "5")

    // Does not wait for command to complete before returning
    err := cmd.Start()
    if err != nil {
        log.Fatal(err)
    }

    // Wait for cmd to Return
    err = cmd.Wait()
    log.Printf("Command finished with error: %v", err)
|======|
