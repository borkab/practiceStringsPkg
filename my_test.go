package my

import (
	"strings"
	"testing"
)

func TestMyTrim(t *testing.T) {
	s := "\n\t\t\t\n Pinke Panke \t\t\n\n\t"
	got := MyTrim(s)
	want := "Pinke Panke"

	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}

func TestMyContains(t *testing.T) {
	t.Run("contains", func(t *testing.T) {
		str := "schmetterling"
		str2 := "met"
		got := MyContains(str, str2)
		want := true

		if got != want {
			t.Errorf("got %#v want %#v", got, want)
		}
	})
	t.Run("doesn't contain", func(t *testing.T) {
		str := "juhairassa"
		str2 := "saa"
		got := MyContains(str, str2)
		want := false

		if got != want {
			t.Errorf("got %#v want %#v", got, want)
		}
	})
}

const str1 = "Pupskanone"

func TestMyHasPrefix(t *testing.T) {

	t.Run("has", func(t *testing.T) {

		str2 := "Pups"
		got := MyHasPrefix(str1, str2)
		want := true

		if got != want {
			t.Errorf("got %#v want %#v", got, want)
		}
	})

	t.Run("doesn't have", func(t *testing.T) {

		str2 := "Furz"
		got := MyHasPrefix(str1, str2)
		want := false

		if got != want {
			t.Errorf("got %#v want %#v", got, want)
		}
	})
}

func TestMyHasSuffix(t *testing.T) {

	t.Run("has", func(t *testing.T) {
		str2 := "ne"
		got := MyHasSuffix(str1, str2)
		want := true

		if got != want {
			t.Errorf("got %#v want %#v", got, want)
		}
	})

	t.Run("has", func(t *testing.T) {
		str2 := "oo"
		got := MyHasSuffix(str1, str2)
		want := false

		if got != want {
			t.Errorf("got %#v want %#v", got, want)
		}
	})
}

func TestMyNewReader(t *testing.T) {
	str := "coding is my cardio"
	got := MyNewReader(str)

	want := strings.NewReader("coding is my cardio")

	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}

/*
[borkab@everest practiceStringsPkg]$ go test
--- FAIL: TestMyNewReader (0.00s)
    my_test.go:98: got &strings.Reader{s:"coding is my cardio", i:0, prevRune:-1} want &strings.Reader{s:"coding is my cardio", i:0, prevRune:-1}
FAIL
exit status 1
FAIL    github.com/borkab/practiceStringsPkg    0.002s
*/
