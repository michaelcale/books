---
Title: Configuring XML parsing and serialization
Id: 332
---

You can modify how a struct is serialized to XML by annotating fields with tags.

By providing `XMLName`, this will be serialized as `<data>...</data>` and not `<ShowXMLName>..</ShowXMLName>`.

@file show_xmlname.go output sha1:5c0022f5314ae77d25d0f51b032b18b68ad5b71a goplayground:xocPt-FbJRj

Fields marked with `xml:"-"` are not serialized.

@file show_omit.go output sha1:a37485c1b814ab4d03bdd3284ca5c068083730cc goplayground:1QgFAqWnqSR

To skip serializing fields with empty values, use `xml:",omitempty"`.

@file show_omit_empty.go output sha1:09048ab7e2b25bf6e77116e0881f6f9eda7a07d2 goplayground:ElorLdu1kcL

To change serialization from XML element to XML attribute use `xml:",attr"`.

@file show_attr.go output sha1:9e7a3e90d3c96f24ee497901a4caf26602bfae6b goplayground:_Y7oX_1RKzp

To change serialization from XML element to character data use `xml:",chardata"`.

@file show_chardata.go output sha1:046c201e766866c557b5d8d6a71b73d203e444da goplayground:D9FI3YXJovI

To change serialization from XML element to CDATA use `xml:",cdata"`.

@file show_cdata.go output sha1:e31576f123acf4c8d78c698361663379654c9ae1 goplayground:zU7hk-kMTpV

To serialize the value verbatim, use `xml:",innerxml"`.

@file show_innerxml.go output sha1:769ee21d974da94b07ee36721b4d668134746a0e goplayground:swii1TyF7qE

To serialize field as XML comment, use `xml:",comment"`. Value can't contain `--`.

@file show_comment.go output sha1:9a7ee3d72c6ce63e74526cfd1950d9d978e81a61 goplayground:CoRf9tZTMJA

`xml:"a>b>str"` nests XML element. This also influences parsing.

@file show_nesting.go output sha1:213ddfaceb3f9582e8e97972aef382878f5a6ba0 goplayground:Xh13bOC50oz
