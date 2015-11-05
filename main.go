package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	testAblePrint(os.Stdout)
}

func testAblePrint(w io.Writer) {
	fmt.Fprintf(w, "test")
}

