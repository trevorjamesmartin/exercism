package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	cows *int
}

func (r *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", *r.cows)
}

func NewSillyNephewError(cows *int) *SillyNephewError {
	return &SillyNephewError{cows: cows}
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if cows < 0 {
		return 0, NewSillyNephewError(&cows)
	}
	fodder, err := weightFodder.FodderAmount()
	if fodder == 0 {
		// nothing to divide
		return 0, err
	}
	switch err {
	case nil:
		if fodder > 0 {
			return fodder / float64(cows), nil
		}
		return 0.0, errors.New("negative fodder")
	case ErrScaleMalfunction:
		if fodder > 0 {
			// double, then divide
			return 2.0 * fodder / float64(cows), nil
		}
		return 0.0, errors.New("negative fodder")
	default:
		return 0.0, err
	}
}
