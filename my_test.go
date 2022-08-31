package my

import "testing"

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
