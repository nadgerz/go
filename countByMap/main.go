// count tallies the number of times each line occurs within a file.
package main

import (
	"fmt"
	"github.com/nadgerz/go/datafile"
	"log"
	"sort"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}

	fmt.Println(counts)

	for name, count := range counts {
		fmt.Printf("%s: %d\n", name, count)
	}

	var names []string
	for name := range counts {
		names = append(names, name)
	}

	fmt.Printf("%#v\n", names)
	sort.Strings(names)
	fmt.Printf("%#v\n", names)

	for _, name := range names {
		fmt.Printf("%s: %d\n", name, counts[name])
	}

	fmt.Printf("%#v\n", counts)
}
