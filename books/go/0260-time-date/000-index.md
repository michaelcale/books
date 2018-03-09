---
Title: Time and date
Id: 166
SOId: 8860
---
The Go [`time`](https://golang.org/pkg/time/) package provides functionality for handling time and date.

Major types in the package:
* structure `time.Time` represents time and date values
* `time.Duration` represents difference between two `time.Time` values in nanoseconds
* `time.Second`, `time.Millisecond` etc. are constants of `time.Duration` that are easier to use than nanoseconds

## Get current time and date

`now := time.Now()`

## Construct a time and date at a given moment in time

`t := time.Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time`

`loc` must be provided and represents time zone. Use `time.UTC` variable for UTC time zone.

## Compare two times for equality

`areEqual := t1.Equal(t2)`

## Add duration to time

`newTime := now.Add(5 * time.Second + time.Millisecond * 100)`

Time values are immutable. `Add` returns a new value.

## Substract duration from time

`newTime := now.Add(-6 * time.Second)`

To substract 6 seconds we add -6 seconds.

## Add years, months, days to time

Adding a duration is for durations smaller than 24 hours.

To advance time by calendar years, months or days, use:

```go
years := 2
months := 3
days := 13
t2 := t.AddDate(years, months, days)
```

## Convert time to Unix representation of time

For interopability with existing code you often need to use Unix time, which is defined as number of seconds since January 1, 1970 UTC.

`unixTime := t.Unix()`

## Get year, month and day from time

`year, month, day := t.Day()`

