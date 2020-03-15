package three_steps_problem_lcci

var (
	preData = make([]int, 1000000+10)
)

func init() {
	preData[1] = 1
	preData[2] = 2
	preData[3] = 4
	for i := 4; i < len(preData); i++ {
		preData[i] = (preData[i-3] + preData[i-2] + preData[i-1]) % 1000000007
	}
}

func waysToStep(n int) int {
	return preData[n]
}
