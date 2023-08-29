package utils

import (
	"strconv"
)

func ToInt64(s interface{}) int64 {
	i, _ := strconv.Atoi(s.(string))
	return int64(i)
}

func ToUint(s interface{}) uint {
	i, _ := strconv.Atoi(s.(string))
	return uint(i)
}

func StoBoll(s interface{}) bool {
	if s.(string) == "true" {
		return true
	}
	return false
}

func ToInt(s interface{}) (int, error) {
	i, err := strconv.Atoi(s.(string))
	if err!=nil {
		return 0,err
	}
	return i,nil
}
