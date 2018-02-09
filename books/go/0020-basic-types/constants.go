package main

// :show start
const (
	i  int = 32       // int constant
	s      = "string" // string constant
	i2     = 33       // untyped number constant

	// this, however, cannot be declared as a constant because []byte is
	// too complicated
	//m []byte = []byte{3, 4}
)

// :show end

func main() {
	// do nothing
}
