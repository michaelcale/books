package main

import "fmt"

// :show start
type Foo struct {
	Bar int
}

func (f Foo) Increment() {
	f.Bar++
}

func main() {
	var p *Foo

	// Calling `p.Increment` is automatically changed to `(*p).Increment` by the compiler.
	// (Note that `*p` is going to remain at 0 because a copy of `*p`, and not the original `*p` are being edited)
	p = &Foo{}
	fmt.Printf("(*p).Bar is %d\n", (*p).Bar)
	p.Increment()
	fmt.Printf("(*p).Bar is %d\n", (*p).Bar)

	// As you can see, calling `(*p).Increment` directly does the same thing.
	p = &Foo{}
	fmt.Printf("(*p).Bar is %d\n", (*p).Bar)
	(*p).Increment()
	fmt.Printf("(*p).Bar is %d\n", (*p).Bar)
}

// :show end
