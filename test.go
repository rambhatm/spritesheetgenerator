package main

import "testing"

func testGetImage(t *testing.T) {
	images := getImageList("sprites")
	if len(images) != 10 {
		t.error
	}
}
