package restoreipaddresses

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := []string{}
	path := []string{}

	validIPNode := func(s string) bool {
		if s == "" {
			return false
		}
		if len(s) > 1 && s[0] == '0' {
			return false
		}
		val, err := strconv.Atoi(s)
		return err == nil && val >= 0 && val <= 255
	}

	var backtracking func(s string, start int)
	backtracking = func(s string, start int) {
		if len(path) == 4 {
			if start == len(s) {
				res = append(res, strings.Join(path, "."))
			}
			return
		}

		for i := start; i < len(s) && i < start+3; i++ {
			str := s[start : i+1]
			if !validIPNode(str) {
				continue
			}

			path = append(path, str)
			backtracking(s, i+1)
			path = path[:len(path)-1]
		}
	}

	if len(s) < 4 || len(s) > 12 {
		return res
	}

	backtracking(s, 0)

	return res
}
