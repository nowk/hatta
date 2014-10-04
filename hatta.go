package hatta

import "net/http"
import "github.com/justinas/alice"
import "github.com/nowk/go-methods"

// New returns a alice.Constructor
func New(mstr string, eh ...http.Handler) alice.Constructor {
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
