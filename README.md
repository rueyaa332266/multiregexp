# multiregexp

Package multiregexp helps to make a set of regular expression. And multiregexp applies functions to logical join the match result in set.

## Installation
```
go get github.com/rueyaa332266/multiregexp
```
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
	"github.com/rueyaa332266/multiregexp"
)

func main() {
	re1 := regexp.MustCompile(`\d`)
	re2 := regexp.MustCompile(`[a-z]`)

    // Make an empty Regexps
    var regs multiregexp.Regexps
    // Append the regexp in Regexps
    regs = multiregexp.Append(regs, re1, re2)

    // Use true to match every regexp
    result1 := regs.MatchString("foo", true)
    fmt.Println(result1)

    // Use false to check if one of the regexp match
    result2 := regs.MatchString("foo", false)
    fmt.Println(result2)

    // Check which regexp in set is matched
    result3 := regs.MatchStringWhich("foo")
    fmt.Println(result3)
}
```