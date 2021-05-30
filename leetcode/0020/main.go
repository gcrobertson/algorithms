package main

import "fmt"

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
*/

func main() {
    ex1()
    ex2()
    ex3()
    ex4()
    ex5()
}

func ex1() {
    input := "()"
    expect := true
    fmt.Printf("Does example 1 work? %v\n", isValid(input) == expect)
}

func ex2() {
    input := "()[]{}"
    expect := true
    fmt.Printf("Does example 2 work? %v\n", isValid(input) == expect)
}

func ex3() {
    input := "(]"
    expect := false
    fmt.Printf("Does example 3 work? %v\n", isValid(input) == expect)
}

func ex4() {
    input := "([)]"
    expect := false
    fmt.Printf("Does example 4 work? %v\n", isValid(input) == expect)
}

func ex5() {
    // input := "{[]}"
    input := "((()(())))"
    expect := true
    fmt.Printf("Does example 5 work? %v\n", isValid(input) == expect)
}

// Node ...
type Node struct {
    val  string
    link *Node
}

// Stack ...
type Stack struct {
    top *Node
}

// Push ...
func (s *Stack) Push(v string) {
    n := &Node{
        val: v,
    }
    if s.top == nil {
        s.top = n
    } else {
        n.link = s.top
        s.top = n
    }
}

// Peek ...
func (s Stack) Peek() *Node {
    if s.top == nil {
        return nil
    }
    return s.top
}

// Pop ...
func (s *Stack) Pop() *Node {
    if s.top == nil {
        return nil
    }
    t := s.top
    s.top = s.top.link
    t.link = nil
    return t
}

// IsEmpty ...
func (s Stack) IsEmpty() bool {
    return s.top == nil
}

// isValid is the algorithm ...
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.3 MB, less than 12.19% of Go online submissions for Valid Parentheses.
func isValid(s string) bool {

    if len(s) == 1 {
        return false
    }

    stack := Stack{
        top: &Node{
            val: string(s[0]),
        },
    }
    s = s[1:]

    // stack.Display()

    for _, t := range s {
        if stack.ValidNext(string(t)) == false {
            return false
        }
    }

    return stack.IsEmpty()
}

// ValidNext ...
func (s *Stack) ValidNext(v string) bool {

    if s.top == nil {
        // return true
        s.Push(v)
        return true
    }

    if s.top.val == "(" && v == ")" || s.top.val == "[" && v == "]" || s.top.val == "{" && v == "}" {
        // just remove the previous and return true. stack will be empty and it catches (), [], {} next to each other ...
        s.Pop()
        return true
    }

    if (s.top.val == "(" || s.top.val == "[" || s.top.val == "{") && (v == "(" || v == "[" || v == "{") {
        // can always open ...
        s.Push(v)
        return true
    }

    return false
}

// Display ...
func (s Stack) Display() {
    fmt.Printf("Display: \t")

    t := s.top

    for t != nil {
        fmt.Printf("%v\t", t.val)
        t = t.link
    }
    fmt.Println()
}