// Runtime: 0 ms, faster than 100.00% of Go online submissions for Defanging an IP Address.
// Memory Usage: 1.9 MB, less than 35.59% of Go online submissions for Defanging an IP Address.
package main

func main() {

}

func defangIPaddr(address string) string {

	var result string

	for _, c := range address {

		if c == '.' {

			result += "[.]"

		} else {

			result += string(c)
		}
	}

	return result

}
