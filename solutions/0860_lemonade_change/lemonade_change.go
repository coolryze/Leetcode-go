package lemonadechange

func lemonadeChange(bills []int) bool {
	fiveCount := 0
	tenCount := 0

	for _, bill := range bills {
		diff := bill - 5
		if diff == 0 {
			fiveCount++
		} else if diff == 5 {
			if fiveCount >= 1 {
				tenCount++
				fiveCount--
			} else {
				return false
			}
		} else if diff == 15 {
			if tenCount >= 1 && fiveCount >= 1 {
				tenCount--
				fiveCount -= 1
			} else if fiveCount >= 3 {
				fiveCount -= 3
			} else {
				return false
			}
		}
	}

	return true
}
