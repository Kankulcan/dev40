package main

import (
	"bytes"
	"fmt"
)

func main() {

}

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
