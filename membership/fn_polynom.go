package membership

import "math"

func Zmf(a, b float64) Function {
	return func(x Value) Value {
		if float64(x) < a {
			return 1
		} else if float64(x) >= a && float64(x) <= b {
			return Value(1 - math.Pow((float64(x)-a)/(b-a), 2))
		} else {
			return 0
		}
	}
}

func Smf(a, b float64) Function {
	return func(x Value) Value {
		if float64(x) < a {
			return 0
		} else if float64(x) >= a && float64(x) <= b {
			return Value(math.Pow((float64(x)-a)/(b-a), 2))
		} else {
			return 1
		}
	}
}

func Pimf(a, b, c, d float64) Function {
	return func(x Value) Value {
		return Smf(a, b)(x) * Zmf(c, d)(x)
	}
}
