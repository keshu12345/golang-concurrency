package main

import (
	"fmt"
	"sync"
)

// List of prefixes
var prefixes = []string{
	// The list of prefixes taken from the images or any other source
	"2y3fEKTS",
	"4VdwEEXC8",
	//... and so on
}

// This is a Trie data structure for efficient prefix matching
type Trie struct {
	children map[rune]*Trie
	end      bool
}

// Insert a word into the Trie
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = &Trie{children: make(map[rune]*Trie)}
		}
		node = node.children[ch]
	}
	node.end = true
}

// Search for the longest prefix in the Trie
func (t *Trie) SearchLongestPrefix(word string) string {
	node := t
	var prefix string
	for _, ch := range word {
		if node.children[ch] != nil {
			prefix += string(ch)
			node = node.children[ch]
		} else {
			break
		}
	}
	return prefix
}

// This function pre-processes the list of prefixes and constructs the Trie in parallel using goroutines
func buildTrie(prefixes []string) *Trie {
	var wg sync.WaitGroup
	root := &Trie{children: make(map[rune]*Trie)}

	for _, prefix := range prefixes {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			root.Insert(p)
		}(prefix)
	}

	wg.Wait()
	return root
}

func main() {
	trie := buildTrie(prefixes)

	// Test
	input := "2y3fEKTSxyz" // Change this to any string to test
	longestPrefix := trie.SearchLongestPrefix(input)
	fmt.Println("Longest Matching Prefix:", longestPrefix)
}
