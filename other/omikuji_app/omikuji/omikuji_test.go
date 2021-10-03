package omikuji_test

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, net/http!")
}

func TestSample(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {t.Fatal("unexpected status code")}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {t.Fatal("unexpected error")}
	const expected = "Hello, net/http!"
	if s := string(b); s != expected {t.Fatalf("unexpected response: %s", s)}
}
