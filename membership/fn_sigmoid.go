package membership

import "math"

func Sigmf(a, b float64) Function {
	return func(x Value) Value {
		return Value(1 / (1 + math.Exp(-a*(float64(x)-b))))
	}
}

func Dsigmf(a1, b1, a2, b2 float64) Function {
	return func(x Value) Value {
		return Sigmf(a1, b1)(x) - Sigmf(a2, b2)(x)
	}
}

func Psigmf(a1, b1, a2, b2 float64) Function {
	return func(x Value) Value {
		return Sigmf(a1, b1)(x) * Sigmf(a2, b2)(x)
	}
}
