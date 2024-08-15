package oneusing

func Max[T int | float64](values []T) (T, error) {

	m := values[0]

	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil

}
