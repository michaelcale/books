---
Title: Use of iota in a bitmask
Id: 9708
Score: 3
---
Iota can be very useful when creating a bitmask. For instance, to represent the state of a network connection which may be secure, authenticated, and/or ready, we might create a bitmask like the following:

    const (
        Secure = 1 << iota // 0b001
        Authn              // 0b010
        Ready              // 0b100
    )

    ConnState := Secure|Authn // 0b011: Connection is secure and authenticated, but not yet Ready

