package pointers

import "testing"

// TestSet tests the Set function.
func TestSet(t *testing.T) {
	x := 15

	Set(&x, 5)

	if x != 5 {
		t.Errorf("want 5, got %d", x)
	}
}

// TestGet tests the Get function.
func TestGet(t *testing.T) {
	x := 15

	v := Get(&x)
	if v != 15 {
		t.Errorf("want 15, got %d", v)
	}

	v = Get[int](nil)
	if v != 0 {
		t.Errorf("want 0, got %d", v)
	}
}
