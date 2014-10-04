package hatta

import "net/http"
import "github.com/nowk/go-methods"

// New returns an alice.Constructor singature middleware bound to a given HTTP
// method. If a custom error handler is not provided, the default handler will
// be used and return a 404 "not found" response when an unallowed method is
// met.
func New(mstr string, eh ...http.Handler) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			m := methods.Allow(mstr)
			if m.Allowed(req) {
				h.ServeHTTP(w, req)

				return
			}

			var e http.Handler
			if len(eh) > 0 {
				e = eh[0]
			} else {
				e = MethodError{}
			}

			e.ServeHTTP(w, req)
		})
	}
}

// MethodError is the default handler used to handle requests that that are not
// allowed
type MethodError struct {
	//\
}

func (m MethodError) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "not found", http.StatusNotFound)
}

// Get shortcut
func Get(eh ...http.Handler) func(http.Handler) http.Handler {
	return New("GET", eh...)
}

// Post shortcut
func Post(eh ...http.Handler) func(http.Handler) http.Handler {
	return New("POST", eh...)
}

// Put shortcut
func Put(eh ...http.Handler) func(http.Handler) http.Handler {
	return New("PUT", eh...)
}

// Patch shortcut
func Patch(eh ...http.Handler) func(http.Handler) http.Handler {
	return New("PATCH", eh...)
}

// Delete shortcut
func Delete(eh ...http.Handler) func(http.Handler) http.Handler {
	return New("DELETE", eh...)
}
