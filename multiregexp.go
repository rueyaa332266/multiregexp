package multiregexp

import (
	"io"
	"regexp"
)

type Regexps struct {
	Regexp []*regexp.Regexp
}

func (res *Regexps) Match(b []byte, matchEvery bool) bool {
	if matchEvery {
		for _, re := range res.Regexp {
			if !re.Match(b) {
				return false
			}
		}
		return true
	} else {
		for _, re := range res.Regexp {
			if re.Match(b) {
				return true
			}
		}
		return false
	}
}

func (res *Regexps) MatchString(s string, matchEvery bool) bool {
	if matchEvery {
		for _, re := range res.Regexp {
			if !re.MatchString(s) {
				return false
			}
		}
		return true
	} else {
		for _, re := range res.Regexp {
			if re.MatchString(s) {
				return true
			}
		}
		return false
	}
}

func (res *Regexps) MatchReader(r io.RuneReader, matchEvery bool) bool {
	if matchEvery {
		for _, re := range res.Regexp {
			if !re.MatchReader(r) {
				return false
			}
		}
		return true
	} else {
		for _, re := range res.Regexp {
			if re.MatchReader(r) {
				return true
			}
		}
		return false
	}
}

func Append(res Regexps, regs ...*regexp.Regexp) Regexps {
	for _, re := range regs {
		res.Regexp = append(res.Regexp, re)
	}
	return res
}
