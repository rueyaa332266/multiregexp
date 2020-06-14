# multiregexp

`multiregexp` is a package for Go that provides the match method for
a list of `regexp`.

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

### Make multiregexp and match

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
    result := regs.MatchString("foo", true)
    fmt.Println(result)

    // Use false to check if one of the regexp match
    result := regs.MatchString("foo", false)
    fmt.Println(result)
}
```