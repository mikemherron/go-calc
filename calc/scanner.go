package calc

import (
	"fmt"
	"strings"
)

type scanner struct {
	tokens   []*token
	position int
}

func errInvalidInput(s string) error {
	return fmt.Errorf("invalid input: %s", s)
}

func newScanner(s string) (*scanner, error) {
	tokens := make([]*token, 0)
	s = stripSpaces(s)
	for len(s) > 0 {
		t := getNextToken(s)
		if t == nil {
			return nil, errInvalidInput(s)
		}
		tokens = append(tokens, t)
		s = s[len(t.value):]
	}

	return &scanner{tokens: tokens}, nil
}

func (s *scanner) next() *token {
	if s.position == len(s.tokens) {
		return &eofToken
	}
	t := s.tokens[s.position]
	s.position++
	return t
}

func (s *scanner) peek() tokenType {
	if s.position == len(s.tokens) {
		return tokenEnd
	}
	return s.tokens[s.position].tokenType
}

func stripSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
