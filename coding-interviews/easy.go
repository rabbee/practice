package coding_interviews

import "strings"

//findRepeatNumber 03. 数组中重复的数字
func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		}
		m[nums[i]] = true
	}
	return nums[0]
}

//findNumberIn2DArray 04. 二维数组中的查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	var i = 0
	var j = len(matrix[i]) - 1
	for i < len(matrix) && j >= 0 {
		if target == matrix[i][j] {
			return true
		}
		if target > matrix[i][j] {
			i++
			continue
		}
		if target < matrix[i][j] {
			j--
			continue
		}
	}
	return false
}

//replaceSpace 05. 替换空格
func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ","%20");
}
