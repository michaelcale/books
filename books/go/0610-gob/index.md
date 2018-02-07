---
Title: gob
Id: 8820
---
## Introduction
Gob is a Go specific serialisation method. it has support for all Go data types except for channels and functions. Gob also encodes the type information into the serialised form, what makes it different from say XML is that it is much more efficient.

The inclusion of type information makes encoding and decoding fairly robust to differences between encoder and decoder.
