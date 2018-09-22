package main

import (
	"fmt"

	"github.com/rabbee/leetcode/medium"
)

func main() {
	var res int
	res = medium.LengthOfLongestSubstring("abcabcbb")
	fmt.Println(res)
	res = medium.LengthOfLongestSubstring("bbbbb")
	fmt.Println(res)
	res = medium.LengthOfLongestSubstring("pwwkew")
	fmt.Println(res)
	res = medium.LengthOfLongestSubstring("尼玛")
	fmt.Println(res)
}
