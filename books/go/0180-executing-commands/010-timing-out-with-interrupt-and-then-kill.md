Title: Timing Out with Interrupt and then Kill
Id: 3521
Score: 2
Body:
    c := exec.Command(name, arg...)
    b := &bytes.Buffer{}
    c.Stdout = b
    c.Stdin = stdin
    if err := c.Start(); err != nil {
        return nil, err
    }
    timedOut := false
    intTimer := time.AfterFunc(timeout, func() {
        log.Printf("Process taking too long. Interrupting: %s %s", name, strings.Join(arg, " "))
        c.Process.Signal(os.Interrupt)
        timedOut = true
    })
    killTimer := time.AfterFunc(timeout*2, func() {
        log.Printf("Process taking too long. Killing: %s %s", name, strings.Join(arg, " "))
        c.Process.Signal(os.Kill)
        timedOut = true
    })
    err := c.Wait()
    intTimer.Stop()
    killTimer.Stop()
    if timedOut {
        log.Print("the process timed out\n")
    }
|======|
