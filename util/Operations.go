package util

func IntDiffAbs(a, b int) int {
	d := a - b
	if d < 0 {
		return -d
	}
	return d
}

func IntAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
