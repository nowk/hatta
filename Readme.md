# hatta

[![Build Status](https://travis-ci.org/nowk/hatta.svg?branch=master)](https://travis-ci.org/nowk/hatta)
[![GoDoc](https://godoc.org/github.com/nowk/hatta?status.svg)](http://godoc.org/github.com/nowk/hatta)

Hatta - Request method check middleware for [alice](https://github.com/justinas/alice)

![The Mad Hatter](https://raw.githubusercontent.com/nowk/hatta/master/MadlHatterByTenniel.jpg)

## Examples

    get := hatta.New("GET")
    chain := alice.New(get, mwa, mwb, ...).Then(handler)

Or

    get := hatta.Get()
    chain := alice.New(get, mwa, mwb, ...).Then(handler)

## License

MIT
    