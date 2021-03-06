package cling // import "github.com/jbowes/cling"

import "github.com/jbowes/cling/skip"

// New returns an error with the supplied text, and a frame from the caller's
// stack.
//
// It is included in this package for completeness.
func New(text string) error { return skip.New(1, text) }

// Errorf returns an error with the supplied format specifier, and a frame from
// the caller's stack.
//
// It is included in this package for completeness.
func Errorf(format string, a ...interface{}) error { return skip.Errorf(1, format, a...) }

// Wrap returns an error wrapping err with the supplied text, and a frame
// from the caller's stack. If err is nil, Wrap returns nil.
//
// The error returned implments the Unwrap method, for programatically
// extracting the error chain.
func Wrap(err error, text string) error { return skip.Wrap(err, 1, text) }

// Wrapf returns an error wrapping err with the supplied format specifier, and a
// frame from the caller's stack.If err is nil, Wrapf returns nil.
//
// The error returned implments the Unwrap method, for programatically
// extracting the error chain.
func Wrapf(err error, format string, a ...interface{}) error { return skip.Wrapf(err, 1, format, a...) }

// Seal returns an error wrapping err with the supplied text, and a frame
// from the caller's stack. If err is nil, Wrap returns nil.
//
// The error returned does not implement Unwrap; the error chain is available
// only for printing.
func Seal(err error, text string) error { return skip.Seal(err, 1, text) }

// Sealf returns an error wrapping err with the supplied format specifier, and a
// frame from the caller's stack. If err is nil, Wrap returns nil.
//
// The error returned does not implement Unwrap; the error chain is available
// only for printing.
func Sealf(err error, format string, a ...interface{}) error { return skip.Sealf(err, 1, format, a...) }
