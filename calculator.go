package calc

func Calculate(s string) (float64, error) {
	scanner, err := newScanner(s)
	if err != nil {
		return 0, err
	}

	exp, err := parse(scanner, 0)
	if err != nil {
		return 0, err
	}

	return exp(), nil
}
