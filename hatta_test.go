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
	return len(w.b), nil
}

func (w *tWriter) Header() http.Header {
	return http.Header{}
}

func (w *tWriter) WriteHeader(n int) {
	//\
}

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

		h := New(v.m)
		req := &http.Request{
			Method: v.m,
		}

		hf := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte(v.b))
		})
		s := h(hf)
		s.ServeHTTP(&buf, req)

		assert.Equal(t, []byte(v.b), buf.b)
	}
}

func TestHattaDefaultError(t *testing.T) {
	var b []byte
	buf := tWriter{b: b}

	h := New("GET")
	req := &http.Request{
		Method: "PUT",
	}

	hf := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("i got put"))
	})
	s := h(hf)
	s.ServeHTTP(&buf, req)

	assert.Equal(t, []byte("not found\n"), buf.b)
}

func TestHattaBYOError(t *testing.T) {
	var b []byte
	buf := tWriter{b: b}

	ef := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("uh-oh"))
	})

	h := New("GET", ef)
	req := &http.Request{
		Method: "PUT",
	}

	hf := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("i got put"))
	})
	s := h(hf)
	s.ServeHTTP(&buf, req)

	assert.Equal(t, []byte("uh-oh"), buf.b)
}

func TestWithAlce(t *testing.T) {
	var b []byte
	buf := tWriter{b: b}

	h := New("POST")
	req := &http.Request{
		Method: "POST",
	}

	hf := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("i got post"))
	})
	chain := alice.New(h).Then(hf)
	chain.ServeHTTP(&buf, req)

	assert.Equal(t, []byte("i got post"), buf.b)
}
