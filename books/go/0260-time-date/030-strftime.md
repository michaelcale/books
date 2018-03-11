---
Title: Format time like strftime
Id: 169
SOId: 80100032
---

If you prefer `strftime` style of formatting time values (as used in Python or Ruby) or porting code that uses that style, you can use one of [several libraries](https://godoc.org/?q=strftime).

Here's an example using `github.com/jehiah/go-strftime` package:

@file strftime.go output noplayground

> Side note on package naming: since the package name is `strftime`, the import path / repository path should be `github.com/jehiah/strftime`. It's bad style to add `go` or `go-` to repository path.

Unfortunately, when it comes to parsing things are not as good. There are two packages:
* https://godoc.org/github.com/jeffjen/datefmt
* https://godoc.org/github.com/knz/strtime

but they both are cgo-wrappers around C libraries, which makes them more finicky to build, especially on Windows.

<!-- TODO: mention https://github.com/bmuller/arrow or write my own library that does both parsing and formatting -->

## List of `strftime` directives

|Code|Meaning|Example|
|----|-------|-------|
|%a|Weekday as locale’s abbreviated name.|Mon|
|%A|Weekday as locale’s full name.|Monday|
|%w|Weekday as a decimal number, where 0 is Sunday and 6 is Saturday.|1|
|%d|Day of the month as a zero-padded decimal number.|30|
|%-d|Day of the month as a decimal number. (Platform specific)|30|
|%b|Month as locale’s abbreviated name.|Sep|
|%B|Month as locale’s full name.|September|
|%m|Month as a zero-padded decimal number.|09|
|%-m|Month as a decimal number. (Platform specific)|9|
|%y|Year without century as a zero-padded decimal number.|13|
|%Y|Year with century as a decimal number.|2013|
|%H|Hour (24-hour clock) as a zero-padded decimal number.|07|
|%-H|Hour (24-hour clock) as a decimal number. (Platform specific)|7|
|%I|Hour (12-hour clock) as a zero-padded decimal number.|07|
|%-I|Hour (12-hour clock) as a decimal number. (Platform specific)|7|
|%p|Locale’s equivalent of either AM or PM.|AM|
|%M|Minute as a zero-padded decimal number.|06|
|%-M|Minute as a decimal number. (Platform specific)|6|
|%S|Second as a zero-padded decimal number.|05|
|%-S|Second as a decimal number. (Platform specific)|5|
|%f|Microsecond as a decimal number, zero-padded on the left.|000000|
|%z|UTC offset in the form +HHMM or -HHMM (empty string if the the object is naive).|
|%Z|Time zone name (empty string if the object is naive).|
|%j|Day of the year as a zero-padded decimal number.|273|
|%-j|Day of the year as a decimal number. (Platform specific)|273|
|%U|Week number of the year (Sunday as the first day of the week) as a zero padded decimal number. All days in a new year preceding the first Sunday are considered to be in week 0.|39|
|%W|Week number of the year (Monday as the first day of the week) as a decimal number. All days in a new year preceding the first Monday are considered to be in week 0.|39|
|%c|Locale’s appropriate date and time representation.|Mon Sep 30 07:06:05 2013|
|%x|Locale’s appropriate date representation.|09/30/13|
|%X|Locale’s appropriate time representation.|07:06:05|
|%%|A literal '%' character.|%|

You can also use http://www.strfti.me/ service to help you build `strftime` formatting strings.
