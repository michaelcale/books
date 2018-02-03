Title: Select and Channels
Id: 3539
Introduction:
The `select` keyword provides an easy method to work with channels and perform more advanced tasks. It is frequently used for a number of purposes:

 - Handling timeouts.
 - When there are multiple channels to read from, the select will randomly read from one channel which has data.
 - Providing an easy way to define what happens if no data is available on a channel. 
|======|
Syntax:
- select {}
- select { case true: }
- select { case incomingData := <-someChannel: }
- select { default: }

|======|
