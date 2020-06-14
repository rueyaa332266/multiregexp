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

func (res *Regexps) MatchWhich(b []byte) []int {
	var match []int
	for i, re := range res.Regexp {
		if re.Match(b) {
			match = append(match, i)
		}
	}
	return match
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

func (res *Regexps) MatchStringWhich(s string) []int {
	var match []int
	for i, re := range res.Regexp {
		if re.MatchString(s) {
			match = append(match, i)
		}
	}
	return match
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

func (res *Regexps) MatchReaderWhich(r io.RuneReader) []int {
	var match []int
	for i, re := range res.Regexp {
		if re.MatchReader(r) {
			match = append(match, i)
		}
	}
	return match
}

func Append(res Regexps, regs ...*regexp.Regexp) Regexps {
	for _, re := range regs {
		res.Regexp = append(res.Regexp, re)
	}
	return res
}
