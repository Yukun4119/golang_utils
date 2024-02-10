package util

func init() {}

func IfThenElse[T any](condition bool, trueValue, falseValue T) T {
    if condition {
        return trueValue
    }
    return falseValue
}