package di

import (
	"bytes"
	"fmt"
)

func Greet(write *bytes.Buffer, name string) {
	fmt.Printf("Hello %s", name)
}
