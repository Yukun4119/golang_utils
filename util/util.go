package util

// IfThenElse is the conditional operator in Golang
// Example:
// x, y := 1, 2
// max := IfThenElse(x > y, x, y).(int)
func IfThenElse[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
