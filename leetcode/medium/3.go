package medium

func LengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		chs := make([]bool, 1000)
		sum := 0
		now := i
		for {
			if now >= len(s) {
				break
			}
			ch := s[now]
			// fmt.Print(string(ch))
			if !chs[ch-'a'] {
				sum++
				chs[ch-'a'] = true
			} else {
				if sum > max {
					max = sum
				}
				break
			}
			if sum > max {
				max = sum
			}
			now++
		}
		// fmt.Println()
	}
	return max
}
