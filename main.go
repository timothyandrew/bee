package main

import (
	"encoding/json"
	"strings"
	"syscall/js"

	"github.com/timothyandrew/bee/search"
)

func jsEntry(d *search.Dictionary, gold rune, nonGold []rune) any {

	// d.Visualize()

	result := search.FindWords(d, search.Options{Gold: gold, NonGold: nonGold})

	b, err := json.Marshal(map[string]interface{}{
		"result": result,
		"err":    nil,
	})
	if err != nil {
		return "{\"err\": \"Failed to marshal JSON\"}"
	}

	return string(b)
}

func main() {
	d := search.NewDictionary()
	err := d.Load()
	if err != nil {
		panic(err)
	}

	js.Global().Set("wasm_entry", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return "{\"err\": \"Incorrect arg count\"}"
		}

		gold := rune(strings.ToLower(args[0].String())[0])

		nonGoldStr := strings.Split(args[1].String(), "")
		nonGold := []rune{}
		for _, s := range nonGoldStr {
			nonGold = append(nonGold, rune(strings.ToLower(s)[0]))
		}

		return jsEntry(d, gold, nonGold)
	}))

	c := make(chan struct{}, 0)
	<-c
}
