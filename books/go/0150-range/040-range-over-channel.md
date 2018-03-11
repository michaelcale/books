---
Title: range over a channel
Id: 110
SOId: 801000o5
---

Iterating over a channel with `range`:

@file range_channel.go output sha1:7de9518cba8b38b2ee3b93750ce36a7555002ba5 goplayground:8b7cqlSMnHb

`range` over a channel receives elements sent over a channel.

Iteration ends when channel is closed.

This is a more convenient version of:

@file range_channel2.go output sha1:e4e103c6e7b3917ca30ad9f0cdc36a4fbb30c618 goplayground:I-H-fNb5ml6
