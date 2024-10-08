package reversewordsinastring

func reverseWords(s string) string {
	strBytes := []byte(s)

	slow := 0
	for fast := 0; fast < len(strBytes); fast++ {
		if strBytes[fast] == ' ' {
			continue
		}

		if slow != 0 {
			strBytes[slow] = ' '
			slow++
		}
		for fast < len(strBytes) && strBytes[fast] != ' ' {
			strBytes[slow] = strBytes[fast]
			fast++
			slow++
		}
	}

	strBytes = strBytes[:slow]
	reverseStr(strBytes[:])

	left := 0
	for right := 0; right <= len(strBytes); right++ {
		if right == len(strBytes) || strBytes[right] == ' ' {
			reverseStr(strBytes[left:right])
			left = right + 1
		}
	}

	return string(strBytes)
}

func reverseStr(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}
