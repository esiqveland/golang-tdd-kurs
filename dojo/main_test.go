package main

import (
	"testing"
)

func TestTennisScore15(t *testing.T) {
	expected := 30
	// call the method under test
	result, _ := addOnePointToThisScore(15, 0)

	if expected != result {
		t.Errorf("error expected '%v' but was '%v'", expected, result)
	}
}

func TestTennisScore0(t *testing.T) {
	expected := 15

	// call the method under test
	result, _ := addOnePointToThisScore(0, 0)

	if expected != result {
		t.Errorf("error expected '%v' but was '%v'", expected, result)
	}

}

func TestTennisScore30(t *testing.T) {
	expected := 40

	// call the method under test
	result, _ := addOnePointToThisScore(30, 0)

	if expected != result {
		t.Errorf("error expected '%v' but was '%v'", expected, result)
	}
}

func TestTennisScore40(t *testing.T) {
	expected := Win

	// call the method under test
	result, _ := addOnePointToThisScore(40, 0)

	if expected != result {
		t.Errorf("error expected '%v' but was '%v'", expected, result)
	}

}

func TestTennisScore0Player2(t *testing.T) {
	expected := 15
	// call the method under test
	_, result := addOnePointToThisScore(0, 0)

	if expected != result {
		t.Errorf("error expected '%v' but was '%v'", expected, result)
	}
}

func TestTennisScore15Player2(t *testing.T) {
	expected := 30
	// call the method under test
	_, result := addOnePointToThisScore(0, 15)

	if expected != result {
		t.Errorf("error expected '%v' but was '%v'", expected, result)
	}
}
