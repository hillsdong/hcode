package trie

import "testing"

func TestTrie(t *testing.T) {
	trie := New()
	trie.Insert("hills")
	trie.Insert("hill")
	trie.Insert("hello")
	trie.Insert("")
	t.Log(trie)

	t.Log(trie.Find("hil"))
	t.Log(trie.Find("hill"))
	t.Log(trie.Find("hills"))
	t.Log(trie.Find(""))

	t.Log(trie.Completion("hi"))
	t.Log(trie.Completion("h"))
	t.Log(trie.Completion("a"))
	t.Log(trie.Completion("hils"))
	t.Log(trie.Completion(""))

	t.Log(trie.Completion2("hi"))
	t.Log(trie.Completion2("h"))
	t.Log(trie.Completion2("a"))
	t.Log(trie.Completion2("hils"))
	t.Log(trie.Completion2(""))

}
