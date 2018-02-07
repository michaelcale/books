---
Title: Logging to syslog
Id: 13468
Score: 0
---
It is also possible to log to syslog with `log/syslog` like this:

```go
package main

import (
    "log"
    "log/syslog"
)

func main() {
    syslogger, err := syslog.New(syslog.LOG_INFO, "syslog_example")
    if err != nil {
        log.Fatalln(err)
    }

    log.SetOutput(syslogger)
    log.Println("Log entry")
}
```

After running we will be able to see that line in syslog:

```text
Jul 26 07:35:21 localhost syslog_example[18358]: 2016/07/26 07:35:21 Log entry
```
