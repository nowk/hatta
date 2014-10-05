package hatta

import "net/http"
import "github.com/nowk/methods"

type MethodCheck struct {
	check *methods.Bouncer
}

// Methods returns a new MethodCheck bound to the provided methods
func Methods(methstr ...string) *MethodCheck {
	c := methods.Allow(methstr...)
	return &MethodCheck{
		check: c,
	}
}

// Else returns an alice.Constructor singature function which errors to the
// provided handler when a request with unallowed method is made.
//
// 		get := Methods("GET")
//		a := get.Else(errHandler)
//
//		chain := alice.New(a, b, c).Then(myHandler)
//
func (m MethodCheck) Else(er http.Handler) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if m.check.Allowed(req) {
				h.ServeHTTP(w, req)
				return
			}

			er.ServeHTTP(w, req)
		})
	}
}

// Get shortcut
func Get() *MethodCheck {
	return Methods("GET")
}

// Post shortcut
func Post() *MethodCheck {
	return Methods("POST")
}

// Put shortcut
func Put() *MethodCheck {
	return Methods("PUT")
}

// Patch shortcut
func Patch() *MethodCheck {
	return Methods("PATCH")
}

// Delete shortcut
func Delete() *MethodCheck {
	return Methods("DELETE")
}
