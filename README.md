<!--
  Attractive html formatting for rendering in github. sorry text editor
  readers! Besides the header and section links, everything should be clean and
  readable.
-->
<h1 align="center">cling</h1>
<p align="center"><i>Clear and obvious error wrapping for <a href="https://golang.org">Go</a> errors</i></p>

<div align="center">
  <a href="https://godoc.org/github.com/jbowes/cling"><img src="https://godoc.org/github.com/jbowes/cling?status.svg" alt="GoDoc"></a>
  <img alt="Alpha Quality" src="https://img.shields.io/badge/status-ALPHA-orange.svg" >
  <a href="https://travis-ci.org/jbowes/cling"><img alt="Build Status" src="https://travis-ci.org/jbowes/cling.svg?branch=master"></a>
  <a href="https://github.com/jbowes/cling/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/release/jbowes/cling.svg"></a>
  <a href="./LICENSE"><img alt="BSD license" src="https://img.shields.io/badge/license-BSD-blue.svg"></a>
  <a href="https://codecov.io/gh/jbowes/cling"><img alt="codecov" src="https://img.shields.io/codecov/c/github/jbowes/cling.svg"></a>
  <a href="https://goreportcard.com/report/github.com/jbowes/cling"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/jbowes/cling"></a>
</div><br /><br />


## Introduction
Introduction | [Examples] | [Contributing] <br /><br />

🚧 ___Disclaimer___: _`cling` is alpha quality software. The API may change
without warning between revisions._ 🚧

`cling` provides a clear and obvious error wrapping API for the new [Go] 2/1.13+
[errors](https://godoc.org/golang.org/x/xerrors) package. If you prefer a
specific function over a [formatting directive](https://godoc.org/golang.org/x/xerrors#Errorf),
and an API that returns a nil error when the error to wrap is nil, then `cling`
is for you.


## Examples
[Introduction] | Examples | [Contributing] <br /><br />

For complete examples and usage, see the [GoDoc documentation](https://godoc.org/github.com/jbowes/cling).

_Wrapping an error_
```go
err := errors.New("an error")
wrapped := cling.Wrap(err, "wrapped")

// Wrapped errors can be programatically inspected
fmt.Print(xerrors.Is(wrapped, err)) // true
```

_Sealing an error_
```go
err := errors.New("an error")
sealed := cling.Wrap(err, "sealed")

// Sealed errors cannot be programatically inspected
fmt.Print(xerrors.Is(sealed, err)) // false
```

Both `Wrap` and `Seal` provide format specifier versions(`Wrapf`, `Sealf`),
as well.

### Building APIs on `cling`

`cling/skip` implements the `cling` API with an additional `skip` argument,
allowing creation of APIs on top of `cling` that will no report themselves in
error caller frames.


## Contributing
[Introduction] | [Examples] | [Usage] | Contributing <br /><br />

I would love your help!

`cling` is still a work in progress. You can help by:

- Trying `oag` against different [OpenAPI] documents, and [reporting bugs][bug]
  when the generated code is broken, or [suggesting improvements][enhancement]
  to the generated code.
- Opening a pull request to resolve an [open issue][issues].
- Adding a feature or enhancement of your own! If it might be big, please
  [open an issue][enhancement] first so we can discuss it.
- Improving this `README` or adding other documentation to `cling`.
- Letting [me] know if you're using `cling`.


<!-- Section link definitions -->
[introduction]: #introduction
[examples]: #examples
[usage]: #usage
[contributing]: #contributing

<!-- Other links -->
[go]: https://golang.org

[issues]: ./issues
[bug]: ./issues/new?labels=bug
[enhancement]: ./issues/new?labels=enhancement

[me]: https://twitter.com/jrbowes
