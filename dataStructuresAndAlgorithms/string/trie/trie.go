package trie

import (
	"bytes"
	"fmt"
)

// New return a new trie tree
func New() *Trie {
	root := &Node{sons: map[byte]*Node{}}
	return &Trie{root}
}

// Node is a node in trie tree
type Node struct {
	data  byte
	isEnd bool
	sons  map[byte]*Node
}

// Trie is a trie tree stuct
type Trie struct {
	root *Node
}

// Insert could insert a word to trie tree
func (t *Trie) Insert(word string) {
	k := len(word)
	if k < 1 {
		return
	}

	n := t.root
	for i := 0; i < k; i++ {
		ns, ok := n.sons[word[i]]
		if !ok {
			ns = &Node{data: word[i], sons: map[byte]*Node{}}
			n.sons[word[i]] = ns
		}
		n = ns
	}
	n.isEnd = true
}

// Find tell you is the word in the trie tree
func (t *Trie) Find(word string) (bool, *Node) {
	k := len(word)
	if k < 1 {
		return false, nil
	}

	n := t.root
	for i := 0; i < k; i++ {
		n = n.sons[word[i]]
		if n == nil {
			return false, nil
		}
	}
	return n.isEnd, n
}

// Completion return a list of words that start with the input word
func (t *Trie) Completion(word string) []string {
	words := []string{}
	_, n := t.Find(word)
	if n == nil {
		return words
	}

	words = t.completion(n, word, words)
	return words
}

func (t *Trie) completion(n *Node, preWord string, words []string) []string {
	for _, ns := range n.sons {
		nsPreWord := preWord + string(ns.data)
		if ns.isEnd {
			words = append(words, nsPreWord)
		}
		words = t.completion(ns, nsPreWord, words)
	}
	return words
}

// Completion2 return a list of words that start with the input word
func (t *Trie) Completion2(word string) []string {
	words := []string{}
	_, n := t.Find(word)
	if n == nil {
		return words
	}

	wordHash := map[*Node]string{n: word}
	toWords := func(n *Node, ns *Node) {
		wordHash[ns] = wordHash[n] + string(ns.data)
		if ns.isEnd {
			words = append(words, wordHash[ns])
		}
		return
	}
	t.completion2(n, toWords)
	return words
}

func (t *Trie) completion2(n *Node, do func(*Node, *Node)) {
	for _, ns := range n.sons {
		do(n, ns)
		t.completion2(ns, do)
	}
	return
}

func (t *Trie) preOrder(n *Node, do func(*Node)) {
	for _, ns := range n.sons {
		do(ns)
		t.preOrder(ns, do)
	}
	return
}

func (t *Trie) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Trie Tree:")

	bufferWriteString := func(n *Node) {
		buffer.WriteString(fmt.Sprintf("%q, ", n.data))
	}

	buffer.WriteString("data = [")
	t.preOrder(t.root, bufferWriteString)
	buffer.WriteString("]")
	return buffer.String()
}
