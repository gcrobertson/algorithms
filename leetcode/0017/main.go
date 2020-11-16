package main

import "fmt"

var digitMap = map[int][]string{
	2: []string{"abc"},
	3: []string{"def"},
	4: []string{"ghi"},
	5: []string{"jkl"},
	6: []string{"mno"},
	7: []string{"pqrs"},
	8: []string{"tuv"},
	9: []string{"wxyz"},
}

// letter case permutation with:
// iteration
// recursion
// trie
// bfs
// dfs

func main() {
	r := letterCasePermutation("23")
	fmt.Printf("result: %v\n", r)
}

func letterCasePermutation(digits string) []string {
	result := []string{}

	if len(digits) == 0 {
		return result
	}

	recursiveLCP()
}

// result = populates result
// digits
// string current combination we are dealing with [?]
// index, where in digits we are
// mapping

func recursiveLCP() {}

/*
// parameter: 	"23"
// return:		[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
func letterCombinations(digits string) []string {

	for _, c := range digits {

		temp := []string
		d, _ := strconv.Atoi(string(c))




		fmt.Printf("%v", d)

		for
	}

}

*/

/*

func digitMap() trie {
    t := newTrie()
    node := t.root
    for i := 0; i < len(digitMap); i++ {
        x := node
        for k := 0; k < len(digitMap[i]); k++ {
            node.children[i] = k
        }
    }
}

func letterCombinations(digits string) []string {

    t := digitMap()

    r := []string{}







    return r
}


type trieNode struct {
    children [26]*trieNode
    isEnd bool
}

type trie struct {
    root trieNode
}

type (t trieNode) insert(digits string) {

}

type (t trieNode) find(digits) []string {

}

func newTrie() trie {
    return trie{
        root : &trieNode{},
    }
}

*/
