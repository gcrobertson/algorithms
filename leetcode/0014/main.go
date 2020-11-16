package main

import "fmt"

type trieNode struct {
	counter  int
	children [26]*trieNode
}

type trie struct {
	root *trieNode
}

func newTrie() trie {
	return trie{
		root: &trieNode{},
	}
}

func (t trie) add(s string) {
	node := t.root
	for i := 0; i < len(s); i++ {
		k := s[i] - 97
		if node.children[k] == nil {
			node.children[k] = &trieNode{}
		}
		node.counter++
		node = node.children[k]
	}
}

func (t *trie) commonPrefix(xs []string) string {

	if len(xs) == 0 {
		return ""
	}

	var result string
	s := xs[2]
	max := t.root.children[s[0]-97].counter
	node := t.root
	for i := 0; i < len(s); i++ {
		k := s[i] - 97

		fmt.Printf("max = %v, node index %v with value %v has count %v\n", max, i, string(s[i]), node.counter)

		if node.children[k].counter < max {
			break
		}
		result += string(s[i])
		node = node.children[k]
	}

	fmt.Printf("the max is %v\n", max)

	return result
}

func main() {

	trie := newTrie()
	fmt.Printf("%T\n", trie)

	xs := []string{"apps", "apples", "appalachia"}
	for _, s := range xs {
		trie.add(s)
	}
	r := trie.commonPrefix(xs)

	fmt.Printf("%v is the common prefix of strings %+v\n", r, xs)
}
