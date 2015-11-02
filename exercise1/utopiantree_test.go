package main

import "testing"

func TestUtopiantree(t *testing.T) {
	result := utopiantree(0)
	if result != 1 {
		t.Fatal("After 0 cycles height should be 1")
	}
	result = utopiantree(1)
	if result != 2 {
		t.Fatal("After 1 cycle height should be 2")
	}
	result = utopiantree(4)
	if result != 7 {
		t.Fatal("After 4 cycles height should be 7")
	}
}
