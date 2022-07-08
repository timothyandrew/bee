package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/timothyandrew/bee/search"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: bee <gold> <non-gold>")
		os.Exit(1)
	}

	gold := rune(strings.ToLower(os.Args[1])[0])

	nonGoldStr := strings.Split(os.Args[2], "")
	nonGold := []rune{}
	for _, s := range nonGoldStr {
		nonGold = append(nonGold, rune(strings.ToLower(s)[0]))
	}

	d := search.NewDictionary()
	err := d.Load()
	if err != nil {
		panic(err)
	}

	// d.Visualize()

	search.FindWords(d, search.Options{Gold: gold, NonGold: nonGold})
}
