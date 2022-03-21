package calc

import (
	"fmt"
	"strings"
)

type scanner struct {
	tokens   []*token
	position int
}

func newScanner(s string) (*scanner, error) {
	tokens := make([]*token, 0)
	s = strings.ReplaceAll(s, " ", "")
	for len(s) > 0 {
		t := getNextToken(s)
		if t == nil {
			panic(fmt.Errorf("no token found for input %s", s))
		}
		tokens = append(tokens, t)
		s = s[len(t.v):]
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

func (s *scanner) peek() *token {
	if s.position >= len(s.tokens)-1 {
		return &eofToken
	}
	return s.tokens[s.position]
}
