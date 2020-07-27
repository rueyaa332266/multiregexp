package multiregexp

import (
	"fmt"
	"regexp"
)

func ExampleRegexps_Match_withAndLogic() {
	var regs Regexps

	re1 := regexp.MustCompile(`\d`)
	re2 := regexp.MustCompile(`[a-z]`)
	regs = Append(regs, re1, re2)
	result := regs.Match([]byte("foo"), "AND")

	fmt.Println(result)
	// Output: false
}

func ExampleRegexps_Match_withOrlogic() {
	var regs Regexps

	re1 := regexp.MustCompile(`\d`)
	re2 := regexp.MustCompile(`[a-z]`)
	regs = Append(regs, re1, re2)
	result := regs.Match([]byte("foo"))

	fmt.Println(result)
	// Output: true
}

func ExampleRegexps_MatchWhich() {
	var regs Regexps

	re1 := regexp.MustCompile(`\d`)
	re2 := regexp.MustCompile(`[a-z]`)
	regs = Append(regs, re1, re2)
	result := regs.MatchWhich([]byte("foo"))

	fmt.Println(result)
	// Output: [1]
}

func ExampleRegexps_MatchString_withAndLogic() {
	var regs Regexps

	re1 := regexp.MustCompile(`\d`)
	re2 := regexp.MustCompile(`[a-z]`)
	regs = Append(regs, re1, re2)
	result := regs.MatchString("foo", "AND")

	fmt.Println(result)
	// Output: false
}

func ExampleRegexps_MatchString_withOrlogic() {
	var regs Regexps

	re1 := regexp.MustCompile(`\d`)
	re2 := regexp.MustCompile(`[a-z]`)
	regs = Append(regs, re1, re2)
	result := regs.MatchString("foo")

	fmt.Println(result)
	// Output: true
}

func ExampleRegexps_MatchStringWhich() {
	var regs Regexps

	re1 := regexp.MustCompile(`\d`)
	re2 := regexp.MustCompile(`[a-z]`)
	regs = Append(regs, re1, re2)
	result := regs.MatchStringWhich("foo")

	fmt.Println(result)
	// Output: [1]
}

func ExampleAppend() {
	var regs Regexps

	re1 := regexp.MustCompile(`\d`)
	re2 := regexp.MustCompile(`[a-z]`)
	regs = Append(regs, re1, re2)

	fmt.Println(regs)
	// Output: [\d [a-z]]
}

func ExampleRegexps_Append() {
	var regs Regexps
	regs1 := Regexps{regexp.MustCompile(`\d`), regexp.MustCompile(`[a-z]`)}
	regs2 := Regexps{regexp.MustCompile(`[0-9]`), regexp.MustCompile(`[A-Z]`)}
	regs = regs.Append(regs1, regs2)

	fmt.Println(regs)
	// Output: [\d [a-z] [0-9] [A-Z]]
}
