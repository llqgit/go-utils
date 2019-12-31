package go_utils

import uuid "github.com/satori/go.uuid"

// 获取一个 string 类型的 UUID
func GetUUID() string {
	var err error
	id := uuid.Must(uuid.NewV4(), err)
	return id.String()
}
