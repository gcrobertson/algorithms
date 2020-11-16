package main

func main() {

}

// Equations are given in the format: A / B = k
// where A and B are variables represented as strings
// and k is a real number (floating point number)
//
//
// Given some queries, return the answers. If the answer does not exist, return -1.0
//
//
// Example:
// Given a / b = 2.0, b / c = 3.0
// Queries are: a / c = ?, b / a ?, a / e ?, a / a ?, x / x?
// Return: [6.0, 0.4, -1.0, 1.0, -1.0]
//
//
//	Equations = [ ["a","b"], ["b","c"] ]
//	Values =	[2.0, 3.0]
//	Queries =	[ ["a","c"], ["b",a"], ["a","e"], ["a","a"], ["x","x"]]
//	Return: 	[6.0, ]
//
//
//
//	a = 2b
//	b = .5a
//	b = 3c
//	a = 6c
//	c = b/3
//  c = a/6
//
//
//	A / B = k
//
//	"a" "b"
//
// equation[0][0] = `k` * equation[0][1]
// equation[1][0] = `k` * equation[1][1]

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

}


sonji.mackey@okta.com 
sonji mackey 
linkedin.com/in/sonji/

1. 1 hour technical screen interview

2. virtual on-site with 4 or 5 people

software engineer 
senior 
staff 		7+ years
principal 	15+ years