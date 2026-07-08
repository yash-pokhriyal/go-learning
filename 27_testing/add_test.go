// Har test file ka naam _test.go se end hona chahiye.

package main

import "testing"

func TestAdd(t *testing.T) {

	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}




// go test ka kaam hai:

// Package ko compile karna
// _test.go files me test functions dhundhna
// Unko run karna
// Batana ki pass hua ya fail


// t.Errorf()

// Ye bolta hai:

// Test fail hua, lekin baaki tests chalne do.


// t.Fatalf()

// Ye bolta hai:

// Test fail hua, ab isi test ko yahin rok do.