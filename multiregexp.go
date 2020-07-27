// Package multiregexp helps to make a set of regular expression. And multiregexp applies functions to logical join the match result in set.
package multiregexp

import (
	"regexp"
)

// Regexps is a set of regular expression.
type Regexps []*regexp.Regexp

// Match reports whether the byte slice b contains any match in the set of the regular expression res.
// When matchType is AND, the result of each match will be logically joined with AND.
// Otherwise the result of each match will be joined with OR.
func (rex Regexps) Match(b []byte, matchType ...string) bool {
	if len(matchType) != 0 && matchType[0] == "AND" {
		for _, re := range rex {
			if !re.Match(b) {
				return false
			}
		}
		return true
	}
	for _, re := range rex {
		if re.Match(b) {
			return true
		}
	}
	return false
}

// MatchWhich reports the index of matched regular expression in the set.
// It returns an empty slice when no regular expression is matched.
func (rex Regexps) MatchWhich(b []byte) []int {
	var match []int
	for i, re := range rex {
		if re.Match(b) {
			match = append(match, i)
		}
	}
	return match
}

// MatchString reports whether the string s contains any match in the set of the regular expression res.
// When matchType is AND, the result of each match will be logically joined with AND.
// Otherwise the result of each match will be joined with OR.
func (rex Regexps) MatchString(s string, matchType ...string) bool {
	if len(matchType) != 0 && matchType[0] == "AND" {
		for _, re := range rex {
			if !re.MatchString(s) {
				return false
			}
		}
		return true
	}
	for _, re := range rex {
		if re.MatchString(s) {
			return true
		}
	}
	return false
}

// MatchStringWhich reports the index of matched regular expression in the set.
// It returns an empty slice when no regular expression is matched.
func (rex Regexps) MatchStringWhich(s string) []int {
	var match []int
	for i, re := range rex {
		if re.MatchString(s) {
			match = append(match, i)
		}
	}
	return match
}

// Append adds regular expression into the set of the regular expression res.
func Append(rex Regexps, regs ...*regexp.Regexp) Regexps {
	for _, re := range regs {
		rex = append(rex, re)
	}
	return rex
}

// Append adds regular expression in Regexps into another Regexps.
func (rex Regexps) Append(rexs ...Regexps) Regexps {
	for _, regexps := range rexs {
		for _, re := range regexps {
			rex = append(rex, re)
		}
	}
	return rex
}
