package medium

import (
	"fmt"
	"testing"
)

func Test1143(t *testing.T) {
	fmt.Println(longestCommonSubsequence("124365", "123564"))
}

func Test576(t *testing.T) {
	fmt.Println(checkInclusion("a", "ab"))
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
	fmt.Println(checkInclusion("hello", "ooollheoooleh"))
	fmt.Println(checkInclusion("hello", "ooolleoooleh"))
}
