---
Title: Format time
Id: 167
SOId: 801000j5
---

Most languages inherited time formatting method from `strftime` function C library which uses somewhat cryptic format strings like `%Y-%m-%d`.

Go designers came up with arguably more intuitive way of time parsing and formatting where you show a template of how you want the result to look like:

@file format.go output sha1:2d56dbcf8e6e21620c5979ec9259843764022f4f goplayground:B_uLdROWGO6

Formatting string is an arbitrary string with some parts being replaced by the data from `time.Time` value:

|template|meaning|
|--------|-------|
|2006, 06|4 or 2 digit year|
|2|month, 1-12|
|1|day, 1-31|
|15|hour|
|am, PM|show hour in am/pm format|
|4|minute|
|5|second|
|MST|string time zone|
|-0700|numeric time zone|
|Jan, January|short or long month name|
|Mon, Monday|short or long day name|

Days, months, years, hours, minutes and seconds can be zero-padded by adding 0 to format number. It'll only be shown for numbers < 10.
`02` means zero-padded month i.e. `04` or `11`.

Some values can be space-padded. `_2` will be ` 4` or `11`.

Package time also defines constants for some well-known formats for date/time formatting e.g. `time.RFC822` is date format defined in RFC 822 which is date format in e-mail messages.

Here's a full list of pre-defined formats:
```go
const (
        ANSIC       = "Mon Jan _2 15:04:05 2006"
        UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
        RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
        RFC822      = "02 Jan 06 15:04 MST"
        RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
        RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
        RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
        RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
        RFC3339     = "2006-01-02T15:04:05Z07:00"
        RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
        Kitchen     = "3:04PM"
        // Handy time stamps.
        Stamp      = "Jan _2 15:04:05"
        StampMilli = "Jan _2 15:04:05.000"
        StampMicro = "Jan _2 15:04:05.000000"
        StampNano  = "Jan _2 15:04:05.000000000"
)
```

What if you prefer `strftime` style of formatting time? That's [available too](169).
