package search

import (
	"fmt"
	"sort"
	"strings"
)

type Options struct {
	Gold    rune
	NonGold []rune
}

func FindWords(d *Dictionary, o Options) []string {
	chars := append(o.NonGold, o.Gold)

	words := FindWordsRec(d, chars, "")

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	result := []string{}
	for _, w := range words {
		if strings.ContainsRune(w, o.Gold) && len(w) >= 4 {
			result = append(result, w)
		}
	}

	return result
}

func FindWordsRec(d *Dictionary, chars []rune, s string) []string {
	values := []string{}

	for _, c := range chars {
		current := fmt.Sprintf("%s%c", s, c)
		node := d.Find(current)

		if node != nil && node.leaf {
			values = append(values, FindWordsRec(d, chars, current)...)
			values = append(values, current)
		} else if node != nil && !node.leaf {
			values = append(values, FindWordsRec(d, chars, current)...)
		}
	}

	return values
}
