package search

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"strings"
)

//go:embed rawdict.txt
var rawDict []byte

// Dictionary trie

type Dictionary struct {
	head *DNode
}

type DNode struct {
	char     rune
	leaf     bool
	children map[rune]*DNode
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		head: &DNode{children: make(map[rune]*DNode)},
	}
}

func (d *Dictionary) Load() error {
	buf := bytes.NewBuffer(rawDict)
	reader := bufio.NewReader(buf)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		d.AddWord(strings.TrimSpace(line))
	}

	return nil
}

func (d *Dictionary) AddWord(word string) {
	current := d.head

	for _, r := range word {
		if node, ok := current.children[r]; ok {
			current = node
		} else {
			node = &DNode{char: r, children: make(map[rune]*DNode)}
			current.children[r] = node
			current = node
		}
	}

	current.leaf = true
}

// Check existence, ignore leaf/non-leaf
func (d *Dictionary) Find(word string) *DNode {
	current := d.head

	for _, r := range word {
		if node, ok := current.children[r]; ok {
			current = node
		} else {
			return nil
		}
	}

	return current
}

func (d *Dictionary) Visualize() {
	d.head.Visualize(0)
}

func (d *DNode) Visualize(level int) {
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}

	if d.leaf {
		fmt.Print("(L)")
	}

	fmt.Println(string(d.char))

	for _, c := range d.children {
		c.Visualize(level + 1)
	}
}
