package c103

func SmallestRangeI1(A []int, K int) int {
	min := 10001
	max := 0
	for i := 0; i < len(A); i++ {
		if A[i] < min {
			min = A[i]
		}
		if max < A[i] {
			max = A[i]
		}
	}
	if K < 0 {
		K = -K
	}
	if min+K >= max-K {
		return 0
	}
	return max - K - min - K
}
