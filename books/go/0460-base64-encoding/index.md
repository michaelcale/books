Title: Base64 Encoding
Id: 4492
Syntax:
- func (enc *base64.Encoding) Encode(dst, src []byte)
- func (enc *base64.Encoding) Decode(dst, src []byte) (n int, err error)
- func (enc *base64.Encoding) EncodeToString(src []byte) string
- func (enc *base64.Encoding) DecodeString(s string) ([]byte, error)
|======|
Remarks:
The [`encoding/base64`](https://godoc.org/encoding/base64) package contains several [built in encoders](https://godoc.org/encoding/base64#pkg-variables). Most of the examples in this document will use `base64.StdEncoding`, but any encoder (`URLEncoding`, `RawStdEncodign`, your own custom encoder, etc.) may be substituted.
|======|
