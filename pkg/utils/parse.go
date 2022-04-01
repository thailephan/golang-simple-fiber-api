package utils

import "strconv"

func ParseInt(value string, bitSize int) (parse int64, err error){
	return strconv.ParseInt(value, 10, bitSize)
}