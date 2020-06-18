// Package multiregexp helps to make a set of regular expression. And multiregexp applies functions to logical join the match result in set.
package multiregexp

import (
	"regexp"
)

// Regexps is a set of regular expression.
type Regexps []*regexp.Regexp

// Match reports whether the byte slice b contains any match in the set of the regular expression res.
// When matchEvery is true, the result of each match will be logically joined with AND.
// Otherwise the result of each match will be joined with OR.
func (res Regexps) Match(b []byte, matchEvery bool) bool {
	if matchEvery {
		for _, re := range res {
			if !re.Match(b) {
				return false
			}
		}
		return true
	} else {
		for _, re := range res {
			if re.Match(b) {
				return true
			}
		}
		return false
	}
}

// MatchWhich reports the index of matched regular expression in the set.
// It returns an empty slice when no regular expression is matched.
func (res Regexps) MatchWhich(b []byte) []int {
	var match []int
	for i, re := range res {
		if re.Match(b) {
			match = append(match, i)
		}
	}
	return match
}

// MatchString reports whether the string s contains any match in the set of the regular expression res.
// When matchEvery is true, the result of each match will be logically joined with AND.
// Otherwise the result of each match will be joined with OR.
func (res Regexps) MatchString(s string, matchEvery bool) bool {
	if matchEvery {
		for _, re := range res {
			if !re.MatchString(s) {
				return false
			}
		}
		return true
	} else {
		for _, re := range res {
			if re.MatchString(s) {
				return true
			}
		}
		return false
	}
}

// MatchStringWhich reports the index of matched regular expression in the set.
// It returns an empty slice when no regular expression is matched.
func (res Regexps) MatchStringWhich(s string) []int {
	var match []int
	for i, re := range res {
		if re.MatchString(s) {
			match = append(match, i)
		}
	}
	return match
}

// Append adds regular expression into the set of the regular expression res.
func Append(res Regexps, regs ...*regexp.Regexp) Regexps {
	for _, re := range regs {
		res = append(res, re)
	}
	return res
}
