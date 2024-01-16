package greetings

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {
	name := "Jagjit"
	want := regexp.MustCompile(`\b`+name+`\b`);
	msg, err := Hello("Jagjit");
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Jagjit") = %q, %v, want match for %#q, nil`, msg, err, want);
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("");
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}