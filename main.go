package main

import(
  "os"
  "fmt"
  "io"
)

func main() {
  testAblePrint(os.Stdout)
}

func testAblePrint(w io.Writer) {
  fmt.Fprintf(w, "test\n")
}
