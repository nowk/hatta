# hatta

[![Build Status](https://travis-ci.org/nowk/hatta.svg?branch=master)](https://travis-ci.org/nowk/hatta)
[![GoDoc](https://godoc.org/github.com/nowk/hatta?status.svg)](http://godoc.org/github.com/nowk/hatta)

Request method check middleware for [alice](https://github.com/justinas/alice)

![The Mad Hatter](https://raw.githubusercontent.com/nowk/hatta/master/MadlHatterByTenniel.jpg)

## Examples

    get := hatta.Methods("GET")
    getCheck := get.Else(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        http.Error(w, "not found", 404)
    }))

    chain := alice.New(getCheck, mwa, mwb, ...).Then(handler)

Can handle multiple method checks

    put := hatta.Methods("PUT", "PATCH")
    putCheck := put.Else(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        http.Error(w, "not found", 404)
    }))

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
    