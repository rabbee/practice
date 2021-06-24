package pkg

import "math"

func Window(origin, target string) (result string) {
	var start, minLen = 0, math.MaxInt64
	var left, right = 0, 0

	var window, needs = make(map[uint8]int), make(map[uint8]int)
	for i := 0; i < len(target); i++ {
		needs[target[i]]++
	}

	var match = 0

	for right < len(origin) {
		addCh := origin[right]
		if _, existed := needs[addCh]; existed {
			window[addCh]++
			if window[addCh] == needs[addCh] {
				match++
			}
		}
		right++

		for match == len(needs) {

			if right-left < minLen {
				start = left
				minLen = right - left
			}

			delCh := origin[left]
			if _, existed := needs[delCh]; existed {
				window[delCh]--
				if window[delCh] < needs[delCh] {
					match--
				}
			}
			left++
		}
	}

	if minLen == math.MaxInt64 {
		return ""
	}
	return origin[start : start+minLen]
}
