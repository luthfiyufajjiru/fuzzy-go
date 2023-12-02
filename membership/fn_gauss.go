package membership

import "math"

func Gaussmf(t, c float64) Function {
	return func(x Value) Value {
		return Value(math.Exp(-math.Pow(float64(x)-c, 2) / (2 * math.Pow(t, 2))))
	}
}

func Gauss2mf(t, c float64) Function {
	return func(x Value) Value {
		return Value(math.Exp(-math.Pow(float64(x)-c, 2) / (2 * math.Pow(t, 2))))
	}
}

func Gbellmf(a, b, c float64) Function {
	return func(x Value) Value {
		return Value(math.Pow(1/(1+math.Abs(float64(x)-c)/a), 2*b))
	}
}
