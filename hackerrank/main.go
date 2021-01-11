package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

/* 	DEMO FUNCTIONALITY:
➜  hackerrank git:(main) ✗ go run main.go
Enter a new password: password   12345
Valid! New password set!
➜  hackerrank git:(main) ✗ go run main.go
Enter your stored password: password   12345
Successful login!

NOTE: Whitespace is getting cut off at the end of the password. For example: "password   12345" stores correctly, "password12345   " is cut off either on write or read...

NOTE: Setting password done in O(n+n+n+n) time complexity could easily be cleaned up into a single loop given more time.

NOTE: Password stored insecurely in plain text. Again, given more time a hash or encryption could be set on write/read from file.

*/

// CLI arguments
var userPassword string
var filepath1 string = "/tmp/password"

func main() {

	pw := readPassword()

	if len(pw) == 0 {
		setPassword()
	} else {
		loginByStoredPassword(pw)
	}

}

// Sets a user password if it has not already been set
func setPassword() bool {
	var valid bool
	reader := bufio.NewReader(os.Stdin)

	for valid != true {
		// Initial user prompt
		fmt.Print("Enter a new password: ")
		// Get user input
		userPassword, _ = reader.ReadString('\n')
		// Validate user input
		valid = validateNewPassword(userPassword)
	}
	storePassword(userPassword)

	return true
}

// Validates a new password against the given rules inlined below
// - at least 10 characters
// - at least 3 whitespace characters
// - at least 1 digit between 4 and 9
// - the following symbols are not allowed: $ ^ & * ( ) [ ]

// @TODO: This could be 1 loop, lack of time to code the problem means I have O(n+n+n+n) time complexity instead ...

func validateNewPassword(password string) bool {

	password = password[0 : len(password)-1]

	// - at least 10 characters
	if len(password) < 10 {
		fmt.Printf("Invalid! Password must be at least 10 characters. Password given is [%v] characters\n", len(password))

		return false
	}

	// - at least 3 whitespace characters
	var count int
	for _, c := range password {
		if unicode.IsSpace(c) {
			count++
		}
	}
	if count < 3 {
		fmt.Printf("Invalid! Password must have at least 3 whitespace characters. Password given has [%v] characters\n", count)

		return false
	}

	// - at least 1 digit between 4 and 9
	var hasNum bool
	for _, c := range password {
		if c > 51 && c < 59 { // going off unicode / ascii range
			hasNum = true
			break
		}
	}
	if hasNum != true {
		fmt.Println("Invalid! Password must have at least 1 digit between 4 and 9")
		return false
	}

	// - the following symbols are not allowed: $ ^ & * ( ) [ ]
	invalidMap := map[int]bool{
		36: true,
		94: true,
		38: true,
		42: true,
		40: true,
		41: true,
		91: true,
		93: true,
	}
	var invalidChars string
	for _, c := range password {
		if _, ok := invalidMap[int(c)]; ok {
			invalidChars += string(c)
		}
	}
	if len(invalidChars) > 0 {
		fmt.Printf("Invalid! Password cannot contain: $ ^ & * ( ) [ ]. Password contains: %s\n", invalidChars)
		return false
	}

	fmt.Println("Valid! New password set!")

	return true
}

// Stores a valid password to a text file
// @TODO: hash or encryption so it is not plain text ...
func storePassword(userPassword string) {

	b := []byte(userPassword)

	err := ioutil.WriteFile(filepath1, b, 0644)
	if err != nil {
		// @TODO: handle error
	}

	// @TODO: Whitespace is getting cut off at the end of the password. For example: "password   12345" stores correctly, "password12345   " is cut off either on write or read...
}

// Reads a stored password
func readPassword() string {
	pw, err := ioutil.ReadFile(filepath1)
	if err != nil {
		// @TODO: error handling
	}

	return string(pw)
}

// Checks user login command line input against stored password
func loginByStoredPassword(pw string) {

	var valid bool
	reader := bufio.NewReader(os.Stdin)

	for valid != true {
		// Initial user prompt
		fmt.Print("Enter your stored password: ")
		// Get user input
		userPassword, _ = reader.ReadString('\n')
		// Validate user input
		valid = (userPassword == pw)
	}

	fmt.Println("Successful login!")
}
