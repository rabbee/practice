package easy_test

import (
	"github.com/rabbee/leetcode/easy"
	"testing"
)

func TestEasy1(t *testing.T) {
	input := []struct {
		Nums   []int
		Target int
		Answer []int
	}{
		{
			Nums:   []int{2, 7, 11, 15},
			Target: 9,
			Answer: []int{0, 1},
		},
		{
			Nums:   []int{2, 6, 11, 15},
			Target: 13,
			Answer: []int{0, 2},
		},
		{
			Nums:   []int{2, 2},
			Target: 4,
			Answer: []int{0, 1},
		},
	}
	for _, item := range input {
		slice := easy.TwoSum(item.Nums, item.Target)
		if len(slice) == 2 && slice[0] == item.Answer[0] && slice[1] == item.Answer[1] {
			continue
		}
		t.Logf("[ERR] nums: %v, target: %d, answer: %v\t| your output: %v\n", item.Nums, item.Target, item.Answer, slice)
		t.Fail()
		break
	}
}
