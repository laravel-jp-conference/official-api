package foundation

import "testing"

func TestFindHallNote(t *testing.T) {
	res := FindHallNote()
	if *res == "" {
		t.Fatalf("it is not expected value %#v", res)
	}
}
