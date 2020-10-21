package main

import "testing"

func TestHello(t *testing.T) {
	resultFromHello := hello()

	if resultFromHello != "Hello, World" {
		t.Errorf("Expected result is Hello, World but %s given", resultFromHello)
	} else {
		t.Logf("Success.")
	}
}
