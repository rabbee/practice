package easy

import (
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
		{
			Nums:   []int{3, 2, 4},
			Target: 6,
			Answer: []int{1, 2},
		},
	}
	for _, item := range input {
		slice := TwoSumV3(item.Nums, item.Target)
		if len(slice) == 2 && slice[0] == item.Answer[0] && slice[1] == item.Answer[1] {
			continue
		}
		t.Logf("[ERR] nums: %v, target: %d, answer: %v\t| your output: %v\n", item.Nums, item.Target, item.Answer, slice)
		t.Fail()
		break
	}
}

func Test121_MaxProfit(t *testing.T) {
	if maxProfit([]int{7, 1, 5, 3, 6, 4}) != 5 {
		t.Fail()
	}
	if maxProfit([]int{7, 6, 4, 3, 1}) != 0 {
		t.Fail()
	}
}
