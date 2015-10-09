package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestPrinter(t *testing.T) {
	expected := "test"
	buf := bytes.NewBuffer([]byte{})

	testAblePrint(buf)

	content, err := ioutil.ReadAll(buf)
	if err != nil {
		t.Errorf("error reading content after print %v", err.Error())
	}
	got := string(content)
	if expected != got {
		t.Errorf("error expected '%+s' but was '%+s'", expected, got)
	}
}
