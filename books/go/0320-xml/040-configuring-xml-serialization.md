---
Title: Configuring XML parsing and serialization
Id: 332
---

The way a struct is serialized as XML can be mofified with struct tags.

By providing `XMLName`, this will be serialized as `<data>...</data>` and not `<ShowXMLName>..</ShowXMLName>`.

@file show_xmlname.go output sha1:874abe0a8ba2d8a13be6e221bf0ef25637b47bd7 goplayground:UqMTIXbHx7D

Fields marked with `xml:"-"` are not serialized.

@file show_omit.go output sha1:cfbe1f35574a7404a1a537bafbce3bccfed24ab7 goplayground:FQnP2kooJQd

To skip serializing fields with empty values, use `xml:",omitempty"`.

@file show_omit_empty.go output sha1:83c3fc08005fb58c8571b00c1b67604f8d913784 goplayground:JQ87g8f7rSG

To change serialization from XML element to XML attribute use `xml:",attr"`.

@file show_attr.go output sha1:ba883fae2c1f5fb264bbf22a442b94b27e062201 goplayground:Snpxe4mDWbS

To change serialization from XML element to character data use `xml:",chardata"`.

@file show_chardata.go output sha1:cc075e220a3e7fedf2ce0ffd51a543c84bbd839e goplayground:zJK6Shib_G5

To change serialization from XML element to CDATA use `xml:",cdata"`.

@file show_cdata.go output sha1:5b5bb63d57d2c86d8063496430339e09012499a6 goplayground:qRRfghR3geg

To serialize the value verbatim, use `xml:",innerxml"`.

@file show_innerxml.go output sha1:d77537333cef56c3f90c6b8f7bc880f497deebe2 goplayground:6UPilA2hEpB

To serialize field as XML comment, use `xml:",comment"`. Value can't contain `--`.

@file show_comment.go output sha1:b85d695eb368302658bf7794fa05d0641499a5f0 goplayground:opaTt0hcMmh

`xml:"a>b>str"` nests XML element.

@file show_nesting.go output sha1:97020fd14c7a33f08ab6134419c21d6ff46539a2 goplayground:G83Dsp9VIXG
