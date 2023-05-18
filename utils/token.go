package utils

import (
	"strconv"
	"strings"
)

func GetNewToken(timeInt int64, userId int) (token string) {
	var build strings.Builder
	build.WriteString(strconv.FormatInt(timeInt, 10))
	build.WriteString(strconv.Itoa(userId))
	build.WriteString(GenValidateCode(6))
	return MD5([]byte(build.String()))
}
