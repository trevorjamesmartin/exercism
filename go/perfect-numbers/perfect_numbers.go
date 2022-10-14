package perfect

import (
	"errors"
)

// Define the Classification type here.
type Classification string

const ClassificationDeficient = "ClassificationDeficient"
const ClassificationAbundant = "ClassificationAbundant"
const ClassificationPerfect = "ClassificationPerfect"

var ErrOnlyPositive = errors.New("only positive integers accepted")

func factors(n int64) []int64 {
	var result []int64
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func aliquotSum(factors []int64) int64 {
	t := int64(0)
	for _, i := range factors {
		t += i
	}
	return t
}

func Classify(n int64) (Classification, error) {
	var nClass Classification
	if n <= 0 {
		return nClass, ErrOnlyPositive
	}
	s := aliquotSum(factors(n))
	switch {
	case s == n:
		nClass = ClassificationPerfect
	case s > n:
		nClass = ClassificationAbundant
	case s < n:
		nClass = ClassificationDeficient
	}
	return nClass, nil
}
