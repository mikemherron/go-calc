package calc

func Calculate(source string) (float64, error) {
	s, err := newScanner(source)
	if err != nil {
		return 0, err
	}

	exp, err := parse(s, 0)
	if err != nil {
		return 0, err
	}

	return exp(), nil
}
