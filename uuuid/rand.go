package uuuid

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	
)

// GenerateUUID 生成UUID并返回其字符串表示形式
func GenerateUUID() (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuidObj.String(), nil
}



func GenerateTimestampID(n int) string {
    randBytes := make([]byte, n/2)
    rand.Read(randBytes)
    return fmt.Sprintf("%x", randBytes)
}