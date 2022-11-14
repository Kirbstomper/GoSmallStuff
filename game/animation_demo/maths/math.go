package maths

func Clamp(i, min, max int) int {
	if i > max {
		return max
	}
	if i < min {
		return min
	}
	return i
}
