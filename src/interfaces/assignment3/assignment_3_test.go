package main

import (
	"github.com/Flaque/filet"
	"testing"
)

type dummyWriter struct {
	out []byte
}

func (s *dummyWriter) Write(p []byte) (n int, err error) {
	(*s).out = p
	return len(p), nil
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestPrintFile(t *testing.T) {
	want := "hello world"
	file := filet.TmpFile(t, "", want)

	d := dummyWriter{}
	PrintFile(file.Name(), &d)

	got := string(d.out)

	if got != want {
		t.Errorf("PrintFile() = %q, want %q", got, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
//func TestHelloEmpty(t *testing.T) {
//	msg, err := Hello("")
//	if msg != "" || err == nil {
//		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
//	}
//}
