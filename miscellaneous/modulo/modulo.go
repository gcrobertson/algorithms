package main

import "fmt"

/*
 *	`Modulo Operation`	:	Finds remainder after division of one number by another
 *							a mod n = modulus
 *
 *	`Modulus`			:	Result of a `Modulo Operation`
 *
 *	`Euclidean Division`:	Dividend ÷ Divisor = Quotient
 *
 *
 *	𝕫					 :   Set of all Natural numbers
 *  ∈					:	"is an element of", denotes set membership.  First letter of Greek word "ἐστί", meaning "is"
 *
 *
 *	`Modulo Operation` satifies:
 *
 *		q ∈ 𝕫
 *		a = nq +r where
 *			- a dividend
 *			- n divisor
 *			- q quotient
 *          - r remainder
 *		|r| < |n|
 */
func main() {
	for i := 0; i < 11; i++ {
		n := 10 - i
		if n == 0 { // panic: runtime error: integer divide by zero for moduloOperation(10, 0)
			break
		}
		fmt.Printf("Result of [%v] %% [%v] = [%v]\n", i, n, moduloOperation(i, n))
	}
	// fmt.Println(moduloOperation(9, 1))  // 0 because there is no remainder
}

/*	a = dividend
 *	n = divisor
 *
 *  Performs Modulo operation returning remainer
 */
func moduloOperation(a int, n int) int {
	return a % n
}
