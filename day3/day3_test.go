package main

import "testing"

func TestGetRingNum(t *testing.T) {
	ring, val := GetRing(1)
	zeroRing := 0
	if ring != zeroRing {
		t.Errorf("Ring number was incorrect, got: %d, want: %d.", ring, zeroRing)
	}
	initialRingVal := 1
	if val != initialRingVal {
		t.Errorf("Ring value was incorrect, got: %d, want: %d.", val, initialRingVal)
	}

	ring, val = GetRing(2)
	firstRing := 1
	if ring != firstRing {
		t.Errorf("Ring number was incorrect, got: %d, want: %d.", ring, firstRing)
	}
	firstRingVal := 9
	if val != firstRingVal {
		t.Errorf("Ring value was incorrect, got: %d, want: %d.", val, firstRingVal)
	}

	ring, val = GetRing(9)
	if ring != firstRing {
		t.Errorf("Ring number was incorrect, got: %d, want: %d.", ring, firstRing)
	}
	if val != firstRingVal {
		t.Errorf("Ring value was incorrect, got: %d, want: %d.", val, firstRingVal)
	}

	ring, val = GetRing(10)
	secondRing := 2
	if ring != secondRing {
		t.Errorf("Ring number was incorrect, got: %d, want: %d.", ring, secondRing)
	}
	secondRingVal := 25
	if val != secondRingVal {
		t.Errorf("Ring value was incorrect, got: %d, want: %d.", val, secondRingVal)
	}
}
