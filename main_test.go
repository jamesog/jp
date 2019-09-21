package main

import (
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func read(file string, t *testing.T) []byte {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func byteDiff(want, got []byte) string {
	return cmp.Diff(string(want), string(got))
}

func TestParseJWT(t *testing.T) {
	token := read("testdata/jwt.io.token", t)
	jwt, err := parseJWT(string(token))
	if err != nil {
		t.Fatal(err)
	}
	header := read("testdata/jwt.io.header", t)
	if diff := byteDiff(header, jwt.header); diff != "" {
		t.Errorf("header mismatch (-want +got):\n%s", diff)
	}
	payload := read("testdata/jwt.io.payload", t)
	if diff := byteDiff(payload, jwt.payload); diff != "" {
		t.Errorf("payload mismatch (-want +got):\n%s", diff)
	}
}
