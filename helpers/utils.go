package helpers

import "strconv"

func StringToInt64(s string) (int64, error) {
	number, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}
	return number, err
}
