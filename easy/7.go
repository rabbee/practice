package easy

const (
	MAX = 1<<31 - 1
	MIN = -(1 << 31)
)

func Reverse(x int) int {
	negative := false
	if x < 0 {
		negative = true
		x = -x
	}
	arrays := make([]int, 0)
	l := 0
	real := false
	for x > 0 {
		tmp := x % 10
		x = x / 10
		if tmp != 0 {
			real = true
		}
		if real {
			arrays = append(arrays, tmp)
		}
	}
	// fmt.Println(arrays)
	for _, dight := range arrays {
		l = l*10 + dight
	}
	if negative {
		l = -l
	}
	if l > MAX || l < MIN {
		l = 0
	}
	return l
}
