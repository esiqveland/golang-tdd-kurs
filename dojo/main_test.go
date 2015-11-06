package main

import (
	"testing"
)

func TestAMethod(t *testing.T) {
	expected := "Does it work?"
	// call the method under test
	result:= methodUnderTest()

	if expected != result {
		t.Errorf("error expected '%v' but was '%v'", expected, result)
	}
}
