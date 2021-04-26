package easy

func TwoSum(nums []int, target int) []int {
	res := []int{
		0,
		1,
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	return res
}

func TwoSumV2(nums []int, target int) []int {
	i, j := 0, 1
	for {
		if nums[i]+nums[j] == target {
			return []int{i, j}
		}
		if j == len(nums)-1 {
			if i < len(nums)-1 {
				i++
				j = i
				if i == len(nums)-1 {
					break
				}
			}
		}
		j++
	}
	return []int{i, j}
}

func TwoSumV3(nums []int, target int) []int {
	intset := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if idx, existed := intset[target-nums[i]]; existed {
			return []int{i, idx}
		}
		intset[nums[i]] = i
	}
	return []int{0, 1}
}
