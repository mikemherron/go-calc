package calc

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

var (
	errMissingClosingBracket = errors.New("missing closing bracket")
)

var bindingPower = map[tokenType]int{
	tokenNumber:       0,
	tokenCloseBracket: 0,
	tokenPlus:         20,
	tokenMinus:        20,
	tokenMultiply:     30,
	tokenDivide:       30,
	tokenPower:        40,
	tokenOpenBracket:  50,
}

func tokenBindingPower(t tokenType) int {
	if p, ok := bindingPower[t]; ok {
		return p
	}
	return 0
}

type expression func() float64

func expressionOpenBracket(s *scanner) (expression, error) {
	inner, err := parse(s, 0)
	if err != nil {
		return nil, err
	}

	next := s.next()
	if next.tokenType != tokenCloseBracket {
		return nil, errMissingClosingBracket
	}

	return inner, nil
}

func expressionNegation(s *scanner) (expression, error) {
	next, err := parse(s, bindingPower[tokenMinus])
	if err != nil {
		return nil, err
	}

	return func() float64 { return next() * -1 }, nil
}

func expressionAddition(left expression, right expression) expression {
	return func() float64 { return left() + right() }
}

func expressionSubtraction(left expression, right expression) expression {
	return func() float64 { return left() - right() }
}

func expressionMultiplication(left expression, right expression) expression {
	return func() float64 { return left() * right() }
}

func expressionDivision(left expression, right expression) expression {
	return func() float64 { return left() / right() }
}

func expressionPower(left expression, right expression) expression {
	return func() float64 { return math.Pow(left(), right()) }
}

func expressionLiteral(t *token) (expression, error) {
	num, err := strconv.ParseFloat(t.value, 64)
	if err != nil {
		return nil, err
	}
	return func() float64 { return num }, nil
}

func parse(s *scanner, rightBindingPower int) (expression, error) {
	left, err := parsePrefixExpression(s.next(), s)
	if err != nil {
		return nil, err
	}

	for tokenBindingPower(s.peek()) > rightBindingPower {
		left, err = parseInfixExpression(left, s.next(), s)
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}

func errInvalidPrefixToken(t *token) error {
	return fmt.Errorf("unable to parse prefix expression for token: %v", t)
}

func parsePrefixExpression(t *token, s *scanner) (expression, error) {
	switch t.tokenType {
	case tokenNumber:
		return expressionLiteral(t)
	case tokenMinus:
		return expressionNegation(s)
	case tokenOpenBracket:
		return expressionOpenBracket(s)
	default:
		return nil, errInvalidPrefixToken(t)
	}
}

func errInvalidInfixToken(t *token) error {
	return fmt.Errorf("unable to parse infix expression for token: %v", t)
}

type infixParser func(expression, expression) expression

func parseInfixExpression(left expression, t *token, s *scanner) (expression, error) {
	var parser infixParser
	switch t.tokenType {
	case tokenPlus:
		parser = expressionAddition
	case tokenMinus:
		parser = expressionSubtraction
	case tokenMultiply:
		parser = expressionMultiplication
	case tokenDivide:
		parser = expressionDivision
	case tokenPower:
		parser = expressionPower
	default:
		return nil, errInvalidInfixToken(t)
	}

	binding := tokenBindingPower(t.tokenType)
	if t.tokenType == tokenPower {
		//Power is right associative
		binding -= 1
	}

	right, err := parse(s, binding)
	if err != nil {
		return nil, err
	}

	return parser(left, right), nil
}
