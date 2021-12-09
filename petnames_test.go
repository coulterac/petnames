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
