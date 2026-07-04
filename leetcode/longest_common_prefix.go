package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	curCursor := -1
	commonPrefix := ""
	for {
		curCursor++

		if curCursor >= len(strs[0]) {
			return commonPrefix
		}

		isAllSame := true
		curChar := strs[0][curCursor]
		for i, _ := range strs {
			if curCursor >= len(strs[i]) {
				isAllSame = false
				break
			}

			if strs[i][curCursor] != curChar {
				isAllSame = false
				break
			}
		}
		if isAllSame {
			commonPrefix += string(curChar)
		} else {
			return commonPrefix
		}
	}

}
