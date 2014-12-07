package hatta

import "net/http"
import "testing"
import "github.com/justinas/alice"
import "github.com/nowk/assert"
import "net/http/httptest"

var err = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "you got error", 400)
})

func TestHattaCallsNextHandlerIfMethodAllowed(t *testing.T) {
	for _, v := range []struct {
		m, b string
	}{
		{"GET", "i got get"},
		{"POST", "i got post"},
		{"PUT", "i got put"},
		{"PATCH", "i got patch"},
		{"DELETE", "i got delete"},
	} {
		w := httptest.NewRecorder()

		m := Methods(v.m)
		check := m.Else(err)

		req := &http.Request{
			Method: v.m,
		}

		hf := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte(v.b))
		})
		s := check(hf)
		s.ServeHTTP(w, req)

		assert.Equal(t, v.b, w.Body.String())
	}
}

func TestHattaHandleError(t *testing.T) {
	w := httptest.NewRecorder()

	get := Get()
	getCheck := get.Else(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "not found", 404)
	}))

	req := &http.Request{
		Method: "PUT",
	}

	hf := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("i got put"))
	})
	s := getCheck(hf)
	s.ServeHTTP(w, req)

	assert.Equal(t, "not found\n", w.Body.String())
}

func TestWithAlice(t *testing.T) {
	w := httptest.NewRecorder()

	post := Post()
	postCheck := post.Else(err)
	req := &http.Request{
		Method: "POST",
	}

	hf := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("i got post"))
	})
	chain := alice.New(postCheck).Then(hf)
	chain.ServeHTTP(w, req)

	assert.Equal(t, "i got post", w.Body.String())
}
