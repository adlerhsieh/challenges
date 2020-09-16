package decode_message

import (
	"strconv"
	"strings"
)

func NumWays(msg string) int {
	return helper(msg, len(msg))
}

func helper(msg string, k int) int {
	if k == 0 {
		return 1
	}
	i := len(msg) - k
	ss := strings.Split(msg, "")
	if ss[i] == "0" {
		return 0
	}

	result := helper(msg, k-1)

	ssZeroAndOne, err := strconv.Atoi(msg[i : i+2])
	if err != nil {
		panic(err)
	}
	if k >= 2 && ssZeroAndOne <= 26 {
		result += helper(msg, k-2)
	}
	return result
}
