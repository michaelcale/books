package main

import "fmt"

func main() {
	// :show start
	const (
		Secure = 1 << iota // 0b001
		Authn              // 0b010
		Ready              // 0b100
	)

	ConnState := Secure | Authn // 0b011: Connection is secure and authenticated, but not yet Ready

	fmt.Printf("Secure: 0x%x (0b%03b)\nAuthn: 0x%x (0b%03b)\nConnState: 0x%x (0b%03b)\n", Secure, Secure, Authn, Authn, ConnState, ConnState)
	// :show end
}
