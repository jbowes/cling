package cling_test

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"

	"github.com/jbowes/cling"
)

func Example() {
	// All funcs in this package print error chains
	err := errors.New("an error")
	wrapped := cling.Wrap(err, "wrapped")
	rewrapped := cling.Wrap(wrapped, "re-wrapped")
	fmt.Printf("%v", rewrapped) // or use %+v to print the stack frames
	// Output: re-wrapped: wrapped: an error
}

func Example_nil() {
	// All funcs in this package return nil when passed a nil error.
	err := error(nil)
	wrapped := cling.Wrap(err, "wrapped")
	fmt.Print(wrapped)
	// Output: <nil>
}

func ExampleWrap() {
	err := errors.New("an error")
	wrapped := cling.Wrap(err, "wrapped")
	fmt.Print(wrapped)
	// Output: wrapped: an error
}

func ExampleWrap_is() {
	err := errors.New("an error")
	wrapped := cling.Wrap(err, "wrapped")

	// Wrapped errors can be programatically inspected
	fmt.Print(xerrors.Is(wrapped, err))
	// Output: true
}

func ExampleWrapf() {
	err := errors.New("an error")
	wrapped := cling.Wrapf(err, "wrapped, especially for %s", "you")
	fmt.Print(wrapped)
	// Output: wrapped, especially for you: an error
}

func ExampleSeal() {
	err := errors.New("an error")
	sealed := cling.Seal(err, "sealed")
	fmt.Print(sealed)
	// Output: sealed: an error
}

func ExampleSeal_is() {
	err := errors.New("an error")
	sealed := cling.Seal(err, "sealed")

	// Sealed errors are opaque for programmatic inspection
	fmt.Print(xerrors.Is(sealed, err))
	// Output: false
}

func ExampleSealf() {
	err := errors.New("an error")
	sealed := cling.Sealf(err, "sealed, especially for %s", "you")
	fmt.Print(sealed)
	// Output: sealed, especially for you: an error
}
