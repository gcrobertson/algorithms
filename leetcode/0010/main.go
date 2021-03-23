package main

func main() {

}

/*
func isMatch(s string, p string) bool {

	// .    matches any single character
	// *    matches zero or more of the
	//        preceding element.

	// s    input string
	// p    pattern

	// ex 1: s = "aa",  p = "a"                     == false
	// ex 2: s = "ab",  p = "a*"                    == true
	// ex 3: s = "ab",  p = ".*"                    == true
	// ex 4: s = "aab", p = "c*a*b"                 == true
	// ex 5: s = "mississippi", p = "mis*is*p*."    == false

	// constrains:
	//   0 <= s.length <= 20
	//   0 <= p.length <= 30

	//  s contains only lowercase English letters.
	//  p contains only lowercase English letters, '.', and '*'.

	var index int

	pattern := p[0]

	for index < len(s) {

		// if s[index] >= 'a' && <= 'z' {

		// }

	}

	for i := 0; i < len(p); i++ {

		if prevChar == "" {
			prevChar = string()
		}

		if p[i] == '.' {
			continue

			// matches any single character
		}

		if p[i] == '*' {

			// * = zero or more of preceding element

			if i > 0 && p[i-1] == '.' {
				// .* means "zero or more (*) of any character (.)
			}

		}

		if i > 0 && p[i-1] == '*' {

		}
	}

}
*/

/*


   if len(p) < len(s) {
       return false    // Handles ex 1.
   }


   idx := map[rune]string {
       '*' : "*",
       '.' : ".",
   }
   // Range over the pattern ...
   for k, c := range p {
       if p[k] == s[k] {
           continue
       }
       if _, ok := idx[c]; ok {
           continue
       }
       return false
   }
   return true
*/
