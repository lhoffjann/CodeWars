package main

import "testing"

func testBSearch(t *testing.T) {
	i := bsearch([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	if i != 2 {
		t.Fatal("seems not to be 2")
	}

}
