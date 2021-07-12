package util

import (
	"strconv"
	"time"
)

type UtcTime time.Time

//实现它的json序列化方法
func (t UtcTime) MarshalJSON() ([]byte, error) {

	b := time.Time(t)

	if b.IsZero() {
		return []byte("0"), nil
	}

	d := TimeMillisecond(b)
	var stamp = strconv.FormatInt(d, 10)
	return []byte(stamp), nil
}

//毫秒转换时间类型
func MillisecondTime(timeUtc int64) time.Time {
	return time.Unix(timeUtc/1000, timeUtc%1000*1e6)
}

//输出毫秒值
func TimeMillisecond(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func TimeStr(timeStr string) (time.Time, error) {
	locationName := "Asia/Shanghai"
	if l, err := time.LoadLocation(locationName); err != nil {
		println(err.Error())
		return time.Time{}, err
	} else {
		lt, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, l)
		return lt, nil
	}
}

func TimeMillisecondStr(str string) int64 {
	timed, _ := TimeStr(str)
	return TimeMillisecond(timed)

}



type TimeSlice []time.Time

func (s TimeSlice) Len() int { return len(s) }

func (s TimeSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s TimeSlice) Less(i, j int) bool { return s[j].Before(s[i]) }

