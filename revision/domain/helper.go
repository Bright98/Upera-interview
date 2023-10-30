package domain

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"time"
)

func LoadEnvFile() error {
	return godotenv.Load(".env")
}
func SetError(key string, message string) *Errors {
	return &Errors{Key: key, Message: message}
}
func GenerateID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}
func NowTime() int64 {
	return time.Now().Unix()
}
