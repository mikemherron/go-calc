package calc

import "testing"

func TestNewScanner(t *testing.T) {

	cases := map[string][]*token{
		"1":      {{tokenNumber, "1"}},
		"10":     {{tokenNumber, "10"}},
		"10.1":   {{tokenNumber, "10.1"}},
		"10.123": {{tokenNumber, "10.123"}},

		"+": {{tokenPlus, "+"}},
		"-": {{tokenMinus, "-"}},
		"*": {{tokenMultiply, "*"}},
		"/": {{tokenDivide, "/"}},

		"10.123 + 10": {
			{tokenNumber, "10.123"},
			{tokenPlus, "+"},
			{tokenNumber, "10"},
		},

		"10.123 + 10 + 1": {
			{tokenNumber, "10.123"},
			{tokenPlus, "+"},
			{tokenNumber, "10"},
			{tokenPlus, "+"},
			{tokenNumber, "1"},
		},
	}

	for input, expected := range cases {
		s, e := newScanner(input)
		if e != nil {
			t.Fatalf(e.Error())
		}
		if !tokensEqual(s.tokens, expected) {
			t.Fatalf("For %s, expected %value to equal %value", input, expected, s.tokens)
		}

	}

}

func tokensEqual(a, b []*token) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if *v != *b[i] {
			return false
		}
	}
	return true
}
