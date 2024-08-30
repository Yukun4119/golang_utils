package util

import "github.com/ryqdev/golang_utils/log"

// IfThenElse is the conditional operator in Golang
// Deprecated: use Ternary instead
// Example:
// x, y := 1, 2
// max := IfThenElse(x > y, x, y).(int)
func IfThenElse[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

// Ternary is the conditional operator in Golang, the updated version of IfThenElse
// Example:
// x, y := 1, 2
// max, ok := Ternary(x > y, x, y).(int)
func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

// DeferRecover
// usage:
// defer DeferRecover()
func DeferRecover() {
	if err := recover(); err != nil {
		log.Error("recover panic: %+v", err)
	}
}
