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

### Build and match

```go
package main

import (
	"fmt"
	"github.com/rueyaa332266/multiregexp"
)

func main() {
	re1 := regexp.MustCompile(`\d`)
	re2 := regexp.MustCompile(`[a-z]`)

    # Build a list of regexp
    var regs multiregexp.Regexps
    regs = multiregexp.Append(regs, re1, re2)

    # Match the regexp

    # Return true when every regexp matched
    result := regs.MatchString("foo", true)
    fmt.Println(result)

    # Return true if one regexp in the list matched
    result := regs.MatchString("foo", false)
    fmt.Println(result)
}
```