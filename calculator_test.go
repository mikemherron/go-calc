package calc

import "testing"

func TestCalculate(t *testing.T) {

	cases := map[string]float64{
		//Basic addition
		"1+1":    2,
		"20+1+4": 25,
		//Basic subtraction
		"20-1":     19,
		"10-12":    -2,
		"100-12-8": 80,
		//Basic multiplication
		"3*9":   27,
		"1*2*3": 6,
		//Basic Division
		"9/3":     3,
		"100/5/4": 5,
		//Basic Combinations
		"20-1+3":   22,
		"20/10*2":  4,
		"20*10/10": 20,
		//Precedence
		"3*2+6":      12,
		"20-1*3":     17,
		"20-1*3+5*2": 27,
		//Parenthesis with no effect
		"(10+1)+10": 21,
		"(3*2)+6":   12,
		//Parenthesis with effect
		"3*(2+6)":                 24,
		"3*((2+6)*2))":            48,
		"(3*((2+6)*2))/(3*(2+6))": 2,
		//Negation
		"-4+2":   -2,
		"-4--2":  -2,
		"-4-2":   -6,
		"-(4*2)": -8,
		// ^ power
		"5^2^3": 390625,
	}

	for input, expected := range cases {
		actual, err := Calculate(input)
		if err != nil {
			t.Errorf("Error on %s: %value", input, err)
		} else if actual != expected {
			t.Errorf("For %s expected %f, got %f", input, expected, actual)
		}
	}
}
