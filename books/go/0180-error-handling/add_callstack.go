package main

import (
	"errors"
	"fmt"

	pkgErrors "github.com/pkg/errors"
)

// :show start

func wrappedError() error {
	err := errors.New("Original error")
	// create a new error value which wraps original err and
	// adds calls tack
	return pkgErrors.WithStack(err)
}

func main() {
	// %+v prints original error add callstack
	fmt.Printf("err: %+v\n\n", wrappedError())

	// errors created with pkg/errors include callstack by default
	fmt.Printf("err: %+v\n", pkgErrors.New("error created with pkg/errors"))
}

// :show end
