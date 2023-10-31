package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"strconv"
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
func GetSkipLimitFromQuery(c *gin.Context) (int64, int64) {
	skip, ok := c.GetQuery("skip")
	if !ok {
		skip = "1"
	}
	limit, ok := c.GetQuery("limit")
	if !ok {
		limit = "10"
	}
	_skip, err := strconv.ParseInt(skip, 10, 64)
	if err != nil {
		_skip = 0
	}
	_limit, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		_limit = 0
	}

	return _skip, _limit
}
