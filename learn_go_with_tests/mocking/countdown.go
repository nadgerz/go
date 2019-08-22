package mocking

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}
