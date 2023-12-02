package constants

import "errors"

var (
	ErrInvalidMembershipFunction = errors.New("invalid membership function")
	ErrNilFuzzySet               = errors.New("nil fuzzy set")
	ErrNilMembershipFunction     = errors.New("nil membership function")
)
