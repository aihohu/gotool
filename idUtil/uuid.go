package idUtil

import (
	"github.com/google/uuid"

	"strings"
)

// RandomUUID 随机UUID
// @return 随机UUID
func RandomUUID() string {
	return uuid.New().String()
}

// SimpleUUID 去掉-的UUID
// @return 去掉-的UUID
func SimpleUUID() string {
	uuidWithHyphen := uuid.New()
	uuidStr := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuidStr
}
