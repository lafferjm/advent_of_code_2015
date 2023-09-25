package box

import "testing"

func TestGetSmallestArea(t *testing.T) {
	b := Box{
		Length: 1,
		Width:  2,
		Height: 3,
	}

	smallest := b.GetSmallestArea()

	if smallest != 2 {
		t.Errorf("Result was incorrect, got: %d, wanted %d", smallest, 2)
	}
}
