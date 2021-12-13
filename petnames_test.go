package petnames

import (
	"testing"
)

func TestName(t *testing.T) {
	pn := new()
	pn2 := new()

	n1 := pn.Name()
	n2 := pn2.Name()

	if n1 == n2 {
		t.Fatalf("expected names to be different")
	}
}

func TestUseDictionary(t *testing.T) {
	dictionary := []string{"frog", "horse", "buffalo"}

	pn := new(
		UseDictionary(dictionary),
	)

	name := pn.Name()

	if ok := contains(dictionary, name); !ok {
		t.Fatalf("expected name to be from provided dictionary")
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
