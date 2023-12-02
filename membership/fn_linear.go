package membership

import (
	"fuzzy-go/constants"

	"github.com/luthfiyufajjiru/errorwrap"
)

func Trimf(a, b, c float64) Function {
	return func(x Value) Value {
		return max(min((x-Value(a))/Value(b-a), (Value(c)-x)/Value(c-b)), 0)
	}
}

func Linsmf(a, b float64) Function {
	return func(x Value) Value {
		if x < Value(a) {
			return 0
		} else if Value(a) <= x && x <= Value(b) && a < b {
			return (x - Value(a)) / Value(b-a)
		} else if x > Value(b) || (x >= Value(b) && a == b) {
			return 1
		}
		panic(errorwrap.New("input out of membership function's range", constants.ErrInvalidMembershipFunction))
	}
}

func Linzmf(a, b float64) Function {
	return func(x Value) Value {
		if x < Value(a) {
			return 1
		} else if Value(a) <= x && x <= Value(b) && a < b {
			return (Value(a) - x) / Value(a-b)
		} else if x > Value(b) || (x >= Value(b) && a == b) {
			return 0
		}
		panic(errorwrap.New("input out of membership function's range", constants.ErrInvalidMembershipFunction))
	}
}

func Trapmf(a, b, c, d float64) Function {
	return func(x Value) Value {
		return max(min((x-Value(a))/Value(b-a), 1, (Value(d)-x)/Value(d-c)), 0)
	}
}
