package world

import "testing"

func TestNew(t *testing.T) {
	w := New()

	if w == nil {
		t.Fatal("expected a World")
	}
}
