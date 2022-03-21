package calc

import (
	"fmt"
	"regexp"
)

type tokenType uint

type tokenMatcher func(string) (bool, string)

func literal(literal string) tokenMatcher {
	return func(s string) (bool, string) {
		return string(s[0]) == literal, literal
	}
}

func pattern(pattern string) tokenMatcher {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		panic(fmt.Errorf("invalid regex expression: %w", err))
	}
	return func(s string) (bool, string) {
		return regex.MatchString(s), regex.FindString(s)
	}
}

const (
	tokenNumber tokenType = iota
	tokenPlus
	tokenMinus
	tokenMultiply
	tokenDivide
	tokenPower
	tokenOpenBracket
	tokenCloseBracket
	tokenEnd
)

var tokenMatchers = map[tokenType]tokenMatcher{
	tokenNumber:       pattern(`^\d+(\.\d+)?`),
	tokenPlus:         literal("+"),
	tokenMinus:        literal("-"),
	tokenMultiply:     literal("*"),
	tokenDivide:       literal("/"),
	tokenOpenBracket:  literal("("),
	tokenCloseBracket: literal(")"),
	tokenPower:        literal("^"),
}

type token struct {
	tokenType tokenType
	v         string
}

func (t *token) String() string {
	return fmt.Sprintf("%v", *t)
}

func getNextToken(s string) *token {
	for t, m := range tokenMatchers {
		if matched, matching := m(s); matched {
			return &token{t, matching}
		}
	}

	return nil
}

var eofToken = token{tokenType: tokenEnd}
