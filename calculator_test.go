package continuousintegration

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)

	if result != 3 {
		t.Errorf("got %v, want %v", result, 3)
	}
}

func TestSubstract(t *testing.T) {
	result := Substract(3, 2)

	if result != 1 {
		t.Errorf("got %v, want %v", result, 1)
	}
}
