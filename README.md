# multiregexp

[![Go Report Card](https://goreportcard.com/badge/github.com/rueyaa332266/multiregexp)](https://goreportcard.com/report/github.com/rueyaa332266/multiregexp)
[![GoDoc](https://godoc.org/github.com/rueyaa332266/multiregexp?status.svg)](https://godoc.org/github.com/rueyaa332266/multiregexp)
![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)

Package multiregexp helps to make a set of regular expression. And it applies functions to logical join the match result in the set.

## Installation
Get the package.
```
go get github.com/rueyaa332266/multiregexp
```

Import in the code.
```
import (
  "github.com/rueyaa332266/multiregexp"
)
```

## Usage

### Make a set of regexp and match the string

```go
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

	// Match every regexp with parameter "AND".
	// re1: not match, re2: match, re3: match. It should be false.
	result1 := regs.MatchString("foo", "AND")
	fmt.Println(result1) // => false

	// Check if one of the regexp match with no second parameter.
	// re1: not match, re2: match, re3: match. It should be true.
	result2 := regs.MatchString("foo")
	fmt.Println(result2) // => true

	// Check which regexp in set is matched
	// re1: not match, re2: match, re3: match. It should get [1 2].
	result3 := regs.MatchStringWhich("foo")
	fmt.Println(result3) // => [1 2]
}
```