package calc

import (
	"errors"
	"fmt"
	"strconv"
)

var bp = map[tokenType]int{
	tokenNumber:       0,
	tokenCloseBracket: 0,
	tokenPlus:         20,
	tokenMinus:        20,
	tokenMultiply:     30,
	tokenDivide:       30,
	tokenPower:        40,
	tokenOpenBracket:  50,
}

func bindingPower(t tokenType) int {
	if p, ok := bp[t]; ok {
		return p
	}
	return 0
}

type expression func() float64

//prefixParser defines a function that takes a token and
//returns an expression
type prefixParser func(*token, *scanner) (expression, error)

//prefixParsers maps token types on to a function that returns a
//prefix expression for that token type.
var prefixParsers map[tokenType]prefixParser

func init() {
	prefixParsers = map[tokenType]prefixParser{
		tokenNumber: func(t *token, s *scanner) (expression, error) {
			num, err := strconv.ParseFloat(t.v, 64)
			if err != nil {
				return nil, err
			}
			return func() float64 { return num }, nil
		},
		tokenOpenBracket: func(t *token, s *scanner) (expression, error) {
			inner, err := parse(s, 0)
			if err != nil {
				return nil, err
			}
			next := s.next()
			if next.tokenType != tokenCloseBracket {
				return nil, errors.New(",missing closing bracket")
			}
			return inner, nil
		},
	}
}

type infixParser func(expression, expression) expression

//infixParsers maps token types on to a function that returns an infix
//expression for that token type
var infixParsers = map[tokenType]infixParser{
	tokenPlus: func(left expression, right expression) expression {
		return func() float64 { return left() + right() }
	},
	tokenMinus: func(left expression, right expression) expression {
		return func() float64 { return left() - right() }
	},
	tokenMultiply: func(left expression, right expression) expression {
		return func() float64 { return left() * right() }
	},
	tokenDivide: func(left expression, right expression) expression {
		return func() float64 { return left() / right() }
	},
}

func parsePrefixExpression(t *token, s *scanner) (expression, error) {
	if _, ok := prefixParsers[t.tokenType]; !ok {
		return nil, errors.New(fmt.Sprintf("no prefix expression available for token type %v", t.tokenType))
	}

	return prefixParsers[t.tokenType](t, s)
}

func parseInfixExpression(left expression, t *token, s *scanner) (expression, error) {
	if _, ok := infixParsers[t.tokenType]; !ok {
		return nil, errors.New(fmt.Sprintf("infixExpression undefined for token type %v", t.tokenType))
	}

	right, err := parse(s, bindingPower(t.tokenType))
	if err != nil {
		return nil, fmt.Errorf("unable to parse right node from %v, %w", t, err)
	}

	return infixParsers[t.tokenType](left, right), nil
}

func parse(s *scanner, rightBindingPower int) (expression, error) {
	left, err := parsePrefixExpression(s.next(), s)
	if err != nil {
		return nil, err
	}

	for bindingPower(s.peek().tokenType) > rightBindingPower {
		left, err = parseInfixExpression(left, s.next(), s)
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}
