# multiregexp

`multiregexp` is a package for Go that provides the match method for
a list of `regexp`.

## Usage

### Building a list of regexp

```go
var regs multiregexp.Regexps

re1 := regexp.MustCompile(`\d`)
re2 := regexp.MustCompile(`[a-z]`)

regs = multierror.Append(regs, re1, re2)
```

### Match with multipleregexp
```go

# Return true when every regexp matched
result := regs.MatchString("foo", true)
fmt.Println(result)

# Return true if one regexp matched
result := regs.MatchString("foo", false)
fmt.Println(result)

```