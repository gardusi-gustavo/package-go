package packagego

import "testing"

func TestAnswer(t *testing.T) {
	if Answer() != 42 {
		t.Fatalf("Answer() = %d", Answer())
	}
}
