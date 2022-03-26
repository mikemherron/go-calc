package calc

import (
	"fmt"
	"regexp"
)

type tokenType uint

func (t tokenType) String() string {
	switch t {
	case tokenNumber:
		return "NUMBER"
	case tokenMinus:
		return "MINUS"
	case tokenMultiply:
		return "MULTIPLY"
	case tokenDivide:
		return "DIVIDE"
	case tokenPower:
		return "POWER"
	case tokenOpenBracket:
		return "OPEN_BRACKET"
	case tokenCloseBracket:
		return "CLOSE_BRACKET"
	case tokenEnd:
		return "EOF"
	}
	return "UNKNOWN"
}

type tokenMatcher func(string) (bool, string)

func literal(literal string) tokenMatcher {
	return func(s string) (bool, string) {
		return string(s[0]) == literal, literal
	}
}

func pattern(pattern string) tokenMatcher {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
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
	value     string
}

func (t *token) String() string {
	return fmt.Sprintf("type:%v, value:%s", t.tokenType, t.value)
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
