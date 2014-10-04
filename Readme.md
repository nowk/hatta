# hatta

[![Build Status](https://travis-ci.org/nowk/hatta.svg?branch=master)](https://travis-ci.org/nowk/hatta)
[![GoDoc](https://godoc.org/github.com/nowk/hatta?status.svg)](http://godoc.org/github.com/nowk/hatta)

Request method check middleware for [alice](https://github.com/justinas/alice)

![The Mad Hatter](https://raw.githubusercontent.com/nowk/hatta/master/MadlHatterByTenniel.jpg)

## Examples

    get := hatta.New("GET")
    chain := alice.New(get, mwa, mwb, ...).Then(handler)

Or

    get := hatta.Get()
    chain := alice.New(get, mwa, mwb, ...).Then(handler)

---

Doesn't directly require [alice](https://github.com/justinas/alice). Returns the `alice.Constructor` signature

    func(http.Handler) http.Handler

Has shortcuts for the basic `HTTP` methods.

    Get()
    Post()
    Put()
    Patch()
    Delete()


## License

MIT
    