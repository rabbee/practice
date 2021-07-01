package medium

func checkInclusion(s1 string, s2 string) bool {
	var left, right int = 0, 0

	fullMap := make(map[uint8]int)
	for i := 0; i < len(s1); i++ {
		fullMap[s1[i]]++
	}

	statMap := make(map[uint8]int)

	var match = 0

	for right < len(s2) {

		if target, existed := fullMap[s2[right]]; existed {
			statMap[s2[right]]++
			if statMap[s2[right]] == target {
				match++
			}
		} else {
			statMap = make(map[uint8]int)
			match = 0
			left = right + 1
		}

		for match == len(fullMap) && left <= right {
			if len(s1) == right-left+1 {
				return true
			}
			if statMap[s2[left]] >= fullMap[s2[left]] {
				statMap[s2[left]]--
			}
			if statMap[s2[left]] < fullMap[s2[left]] {
				match--
			}
			left++
		}

		right++
	}

	return false
}
