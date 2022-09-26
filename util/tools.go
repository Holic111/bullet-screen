package util

import (
	"github.com/google/uuid"
	"strconv"
	"strings"
)

// 拼接字符串
func ConsistString(strs ...string) string {
	sb := strings.Builder{}

	for _, v := range strs {
		sb.WriteString(v)
	}

	return sb.String()
}

// 将字符串转换为整数
func StringToInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return ERROR, err
	}
	return num, nil
}

// 获取随机的UUID
func GetUUID() string {
	newUUID, _ := uuid.NewUUID()

	return newUUID.String()
}