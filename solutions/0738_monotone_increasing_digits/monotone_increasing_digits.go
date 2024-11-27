package monotoneincreasingdigits

import "strconv"

func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	ss := []byte(s)
	flag := len(ss)

	for i := len(ss) - 1; i > 0; i-- {
		if ss[i-1] > ss[i] {
			flag = i
			ss[i-1] -= 1
		}
	}
	for i := flag; i < len(ss); i++ {
		ss[i] = '9'
	}

	res, _ := strconv.Atoi(string(ss))
	return res
}
