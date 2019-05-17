package skip

import (
	"fmt"

	"golang.org/x/xerrors"
)

// New returns an error with the supplied text, and a frame from the caller's
// stack. The argument skip is the number of frames to skip over. Caller(0)
// returns the frame for the caller of New.
//
// This func is intended to be used for implementing APIs on top of cling.
func New(skip uint, text string) error {
	return &sealError{
		msg:   text,
		frame: xerrors.Caller(int(skip + 1)),
	}
}

// Errorf returns an error with the supplied format specifier, and a frame
// from the caller's stack. The argument skip is the number of frames to skip
// over. Caller(0) returns the frame for the caller of Errorf.
//
// This func is intended to be used for implementing APIs on top of cling.
func Errorf(skip uint, format string, a ...interface{}) error {
	return &sealError{
		msg:   fmt.Sprintf(format, a...),
		frame: xerrors.Caller(int(skip + 1)),
	}
}

// Wrap returns an error wrapping err with the supplied text, and a frame
// from the caller's stack. The argument skip is the number of frames to skip
// over. Caller(0) returns the frame for the caller of Wrap. If err is nil, Wrap
// returns nil.
//
// The error returned implments the Unwrap method, for programatically
// extracting the error chain.
//
// This func is intended to be used for implementing APIs on top of cling.
func Wrap(err error, skip uint, text string) error {
	if err == nil {
		return nil
	}

	return &wrapError{sealError{
		msg:   text,
		err:   err,
		frame: xerrors.Caller(int(skip + 1)),
	}}
}

// Wrapf returns an error wrapping err with the supplied format specifier, and a
// frame from the caller's stack. The argument skip is the number of frames to
// skip over. Caller(0) returns the frame for the caller of Wrap. If err is nil,
// Wrapf returns nil.
//
// The error returned implments the Unwrap method, for programatically
// extracting the error chain.
//
// This func is intended to be used for implementing APIs on top of cling.
func Wrapf(err error, skip uint, format string, a ...interface{}) error {
	if err == nil {
		return nil
	}

	return &wrapError{sealError{
		msg:   fmt.Sprintf(format, a...),
		err:   err,
		frame: xerrors.Caller(int(skip + 1)),
	}}
}

// Seal returns an error wrapping err with the supplied text, and a frame
// from the caller's stack. The argument skip is the number of frames to skip
// over. Caller(0) returns the frame for the caller of Wrap. If err is nil, Wrap
// returns nil.
//
// The error returned does not implement Unwrap; the error chain is available
// only for printing.
//
// This func is intended to be used for implementing APIs on top of cling.
func Seal(err error, skip uint, text string) error {
	if err == nil {
		return nil
	}

	return &sealError{
		msg:   text,
		err:   err,
		frame: xerrors.Caller(int(skip + 1)),
	}
}

// Sealf returns an error wrapping err with the supplied format specifier, and a
// frame from the caller's stack. The argument skip is the number of frames to
// skip over. Caller(0) returns the frame for the caller of Wrap. If err is nil,
// Wrap returns nil.
//
// The error returned does not implement Unwrap; the error chain is available
// only for printing.
//
// This func is intended to be used for implementing APIs on top of cling.
func Sealf(err error, skip uint, format string, a ...interface{}) error {
	if err == nil {
		return nil
	}

	return &sealError{
		msg:   fmt.Sprintf(format, a...),
		err:   err,
		frame: xerrors.Caller(int(skip + 1)),
	}
}

type sealError struct {
	msg   string
	err   error
	frame xerrors.Frame
}

func (e *sealError) Error() string              { return fmt.Sprint(e) }
func (e *sealError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

func (e *sealError) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.msg)
	e.frame.Format(p)
	return e.err
}

type wrapError struct {
	sealError
}

func (e *wrapError) Unwrap() error { return e.err }
