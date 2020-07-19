package main

import (
	"fmt"
	"regexp"

  "github.com/rueyaa332266/multiregexp"
)

func main() {
	// Make an empty Regexps
	var regs multiregexp.Regexps

	re1 := regexp.MustCompile(`\d`)
	re2 := regexp.MustCompile(`[a-z]`)
	re3 := regexp.MustCompile(`foo`)

	// Append the regexp in Regexps
	regs = multiregexp.Append(regs, re1, re2, re3)

	// Use true to match every regexp
	result1 := regs.MatchString("foo", true)
	fmt.Println(result1) // => false

	// Use false to check if one of the regexp match
	result2 := regs.MatchString("foo", false)
	fmt.Println(result2) // => true

	// Check which regexp in set is matched
	result3 := regs.MatchStringWhich("foo")
	fmt.Println(result3) // => [1 2]
}