# hatta

[![Build Status](https://travis-ci.org/nowk/hatta.svg?branch=master)](https://travis-ci.org/nowk/hatta)
[![GoDoc](https://godoc.org/github.com/nowk/hatta?status.svg)](http://godoc.org/github.com/nowk/hatta)

Hatta - Request method check middleware for [alice](https://github.com/justinas/alice)

## Examples

    h := hatta.New("GET")
    chain := alice.New(h, mwa, mwb, ...).Then(handler)

## License

MIT
    