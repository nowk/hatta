package hatta

import "net/http"
import "testing"
import "github.com/justinas/alice"
import "github.com/nowk/assert"

type tWriter struct {
	b []byte
}

func (w *tWriter) Write(b []byte) (int, error) {
	w.b = append(w.b, b...)
	return len(b), nil
}

func (w *tWriter) Header() http.Header {
	return http.Header{}
}

func (w *tWriter) WriteHeader(n int) {
	//\
}

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
		var b []byte
		buf := tWriter{b: b}

		m := Methods(v.m)
		check := m.Else(err)

		req := &http.Request{
			Method: v.m,
		}

		hf := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte(v.b))
		})
		s := check(hf)
		s.ServeHTTP(&buf, req)

		assert.Equal(t, []byte(v.b), buf.b)
	}
}

func TestHattaHandleError(t *testing.T) {
	var b []byte
	buf := tWriter{b: b}

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
	s.ServeHTTP(&buf, req)

	assert.Equal(t, []byte("not found\n"), buf.b)
}

func TestWithAlice(t *testing.T) {
	var b []byte
	buf := tWriter{b: b}

	post := Post()
	postCheck := post.Else(err)
	req := &http.Request{
		Method: "POST",
	}

	hf := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("i got post"))
	})
	chain := alice.New(postCheck).Then(hf)
	chain.ServeHTTP(&buf, req)

	assert.Equal(t, []byte("i got post"), buf.b)
}
