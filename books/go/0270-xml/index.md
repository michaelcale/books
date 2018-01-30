Title: XML
Id: 1846
Remarks:
While many uses of the [`encoding/xml`](https://godoc.org/encoding/xml) package include marshaling and unmarshaling to a Go `struct`, it's worth noting that this is not a direct mapping. The package documentation states:

> Mapping between XML elements and data structures is inherently flawed:
> an XML element is an order-dependent collection of anonymous values,
> while a data structure is an order-independent collection of named values.

For simple, unordered, key-value pairs, using a different encoding such as Gob's or [JSON](http://stackoverflow.com/documentation/go/994/json#t=201607261844536104265) may be a better fit. For ordered data or event / callback based streams of data, XML may be the best choice.
|======|
