package mymath

func SumFloat64(s []float64) float64 {
	r := 0.0
	for _, v := range s {
		r += v
	}
	return r
}

func CumSumFloat(s []float64) []float64 {
	if len(s) <= 1 {
		return s
	}
	r := make([]float64, len(s))
	r[0] = s[0]
	for i := 1; i < len(s); i++ {
		r[i] = r[i-1] + s[i]
	}
	return r
}

func FindFloat64(s []float64, r float64) int {
	for k, v := range s {
		if v >= r {
			return k
		}
	}
	return len(s)
}

func DelInt(s []int, index int) []int {
	for k, v := range s {
		if v == index {
			s = append(s[:k], s[k+1:]...)
			return s
		}
	}
	return []int{}
}
