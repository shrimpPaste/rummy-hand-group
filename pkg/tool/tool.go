package tool

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"rummy-logic-v3/global"
	"strconv"
	"time"
)

func GenerateRandomString(length int) string {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		global.Logger.Error(context.Background(), "生成随机字符串错误", err)
		return ""
	}

	randomString := uuidObj.String()[:length]
	return randomString
}

func getStartOfDayTimestamp(timestamp int64) int64 {
	t := time.Unix(timestamp, 0) // 根据时间戳创建 time.Time 对象
	startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return startOfDay.Unix()
}

// GetStartAndEndOfDayTimestamps 获取指定时间戳当天零点和最后一秒的时间戳
func GetStartAndEndOfDayTimestamps(timestamp uint) (uint, uint) {
	t := time.Unix(int64(timestamp), 0) // 根据时间戳创建 time.Time 对象

	// 获取当天零点的时间对象
	startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	startOfDayTimestamp := startOfDay.Unix()

	// 获取当天最后一秒的时间对象
	endOfDay := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	endOfDayTimestamp := endOfDay.Unix()

	return uint(startOfDayTimestamp), uint(endOfDayTimestamp)
}

func Divide(a, b float64) float64 {
	if b == 0 {
		return a
	}
	return a / b
}

func Float64ToString(data float64) string {
	if data == 0 {
		return "0.00%"
	}
	return fmt.Sprintf("%.2f", data*100) + "%"
}

func StringToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println("字符串转float64错误", err.Error())
		return 0
	}
	return f
}

// GenerateSnowflakeID 生成唯一的雪花 ID 并返回字符串类型
func GenerateSnowflakeID() string {
	timestamp := time.Now().UnixNano()
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63()
	snowflakeID := (timestamp << 22) | (randomNumber & 0x3FFFFF)
	if snowflakeID < 0 {
		snowflakeID = -snowflakeID
	}
	return fmt.Sprintf("%d", snowflakeID)
}
