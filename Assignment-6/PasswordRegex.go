/*Number of characters should be 7 or more
● Should have at least 1 number, atleast 1 UpperCase and atleast 1
lowercase letter and 1 special character limited to 3 values( _, @,- )*/

package main

import (
	"fmt"
	"regexp"
)

func main() {

	pattern := "^(=?.*[0-9])(=?.*[A-Z])(=?.*[a-z])(=?.*[_@-])[A-Za-z0-9_@-]{7,}$"
	var password string
	fmt.Println("Enter Password:")
	fmt.Scanln(&password)

	regexPatternCompiled, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error while compiling regex pattern", err)
		return

	}
	matches := regexPatternCompiled.MatchString(password)
	fmt.Println(matches)
}
