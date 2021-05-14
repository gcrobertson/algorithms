package main 

import "fmt"

func ex1() {
	x := 2.00000
	n := 10
	e := 1024.00000
	fmt.Printf("Ex 1 is correct? %v\n", myPow(x, n) == e)
}

func ex2() {
	x := 2.10000
	n := 3
	e := 9.261000000000001
	fmt.Printf("Ex 2 is correct? %v\n", myPow(x, n) == e)
}

func ex3() {
	x := 2.00000
	n := -2
	e := 0.25000
	fmt.Printf("Ex 3 is correct? %v\n", myPow(x, n) == e)
}


func main() {

	ex1()
	ex2()
	ex3()

}

func myPow(x float64, n int) float64 {
    
   
    if n < 0 {
        x = 1 / x
        n = -n 
    }
    
    return fastPow(x, n)
}

func fastPow(x float64, n int) float64 {
    
    if n == 0 {
        return 1.0
    }
    
    half := fastPow(x, n/2)
    
    if n % 2 == 0 {
        return half * half
    } else {
        return half * half * x
    }
}

/*
func myPow(x float64, n int) float64 {
    
   
    if n < 0 {
        x = 1 / x
        n = -n 
    }
    
    ans := 1.0
    cur := x
    
    for i := n; i > 0; i /= 2 {
        
        if i % 2 == 1 {
            ans = ans * cur
        }
        
        cur = cur * cur
    }
    
    return ans
}
*/