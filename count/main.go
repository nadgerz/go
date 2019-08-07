// count tallies the number of times each line occurs within a file.
package main

import (
	"fmt"
	"github.com/nadgerz/go/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines)
}
