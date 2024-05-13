package vanity_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"jordanreger.com/vanity"
)

var config string

func init() {
	b, _ := os.ReadFile("vanity.json")
	config = string(b)
}

func TestExample1(t *testing.T) {
	mux := http.NewServeMux()

	vanity.Serve(mux, config)

	server := httptest.NewServer(mux)
	defer server.Close()

	res, err := http.Get(server.URL + "/example1/")
	if err != nil {
		t.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(body))
}

func TestExample2(t *testing.T) {
	mux := http.NewServeMux()

	vanity.Serve(mux, config)

	server := httptest.NewServer(mux)
	defer server.Close()

	res, err := http.Get(server.URL + "/example2/")
	if err != nil {
		t.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(body))
}
