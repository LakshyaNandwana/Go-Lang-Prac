package oneusinginterface

type Ordered interface {
	~int | ~float64
}

func Max[T Ordered](values []T) (T, error) {

	m := values[0]

	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil

}
