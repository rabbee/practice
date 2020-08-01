package hard

//76. 最小覆盖子串
//给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。
//
//示例：
//
//输入: S = "ADOBECODEBANC", T = "ABC"
//输出: "BANC"
//说明：
//
//如果 S 中不存这样的子串，则返回空字符串 ""。
//如果 S 中存在这样的子串，我们保证它是唯一的答案。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-window-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func minWindow(s string, t string) string {
	var lp, rp = 0, 0

	fullSet := make(map[uint8]bool)
	targetCountMap := make(map[uint8]int)
	countMap := make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		if _, inMap := targetCountMap[t[i]]; inMap {
			targetCountMap[t[i]]++
		} else {
			targetCountMap[t[i]] = 1
		}
	}

	var bestLp, bestRp = 0, len(s)
	//fmt.Println(targetCountMap)
	//fmt.Println(countMap)

	for {
		ch := s[rp]
		//fmt.Println(countMap)
		if targetCount, ok := targetCountMap[ch]; ok {
			if _, inMap := countMap[ch]; inMap {
				countMap[ch]++
			} else {
				countMap[ch] = 1
			}
			if countMap[ch] == targetCount {
				fullSet[ch] = true
			}
			//fmt.Printf("%d, %d\n", len(fullSet), len(targetCountMap))
			if len(fullSet) == len(targetCountMap) { //已覆盖字串
				isSkip := false
				for lp < rp { //过滤无关字符
					lch := s[lp]
					//fmt.Sprintf("%d, %d\n", lp, rp)
					if targetCount, ok := targetCountMap[lch]; ok {
						if count, inMap := countMap[lch]; inMap {
							if count-1 < targetCount { //删除后不覆盖
								if bestRp-bestLp > rp-lp {
									bestLp = lp
									bestRp = rp
								}
								delete(fullSet, lch)
								countMap[lch]--
								lp++
								isSkip = true
								break
							}
							countMap[lch]--
						}
					}
					lp++
				}
				if !isSkip {
					if bestRp-bestLp > rp-lp {
						bestLp = lp
						bestRp = rp
					}
				}
			}
		}
		rp++
		if rp == len(s) {
			break
		}
	}
	if bestLp == 0 && bestRp == len(s) {
		return ""
	}
	return s[bestLp : bestRp+1]
}
