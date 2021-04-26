package easy

func reverse(x int) int {
	n := x
	if x < 0 {
		n = -n
	}
	y := 0
	for ; n > 0; n /= 10 {
		if y > MAX/10 {
			return 0
		}
		y = y*10 + n%10
	}
	if x < 0 {
		return -y
	}
	return y
}
