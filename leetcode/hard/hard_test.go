package hard

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T)  {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("aa", "aa"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}
