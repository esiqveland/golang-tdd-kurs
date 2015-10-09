package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestPrinter(t *testing.T) {
	expected := []byte("test\n")
	buf := bytes.NewBuffer([]byte{})

	testAblePrint(buf)

	content, err := ioutil.ReadAll(buf)
	if err != nil {
		t.Errorf("error reading content after print %v", err.Error())
	}
	if string(expected) != string(content) {
		t.Errorf("error expected '%+v' but was '%+v'", expected, content)
	}
}
