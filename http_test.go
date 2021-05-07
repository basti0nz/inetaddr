package inetaddr

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpCheck(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "127.0.0.1")
	}))
	defer ts.Close()
	ts2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, " 127.0.0.1 ")
	}))
	ts3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, " hello, world! ")
	}))
	ts4 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}))

	defer ts.Close()
	defer ts2.Close()
	defer ts3.Close()
	defer ts4.Close()

	httpTestServerS := []HttpServer{
		{ts.URL},
		{ts2.URL},
	}

	for i := range httpTestServerS {
		ip, err := HttpCheck(httpTestServerS[i])
		if err != nil {
			t.Errorf("Expected IP but has  %s", err)
		}
		assert.Equal(t, "127.0.0.1", ip.String())

	}

	_, err := HttpCheck(HttpServer{ts3.URL})
	if err != nil {
		assert.Equal(t, fmt.Sprint(err), "Response error: not ip address hello, world!")
	}

	_, err = HttpCheck(HttpServer{ts4.URL})
	if err != nil {
		assert.Equal(t, fmt.Sprint(err), "Response error: status code is 404")
	}

}
