package util

import (
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func StringToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func Int64UUID() int64 {
	d := rand.Int63n(999)
	t := time.Now().UnixNano()
	return d + t
}
