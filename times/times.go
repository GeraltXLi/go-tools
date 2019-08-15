//Package tims provides functionality for getting particular timestamp or date.
//Time format:2006-01-02 03:04:05, 2006-01-02 15:04:05, etc.

package times

import (
	"time"
)

//var cstZone = time.FixedZone("CST",8*3600)
var (
	cstZone       *time.Location
	LocationTimes int64
)

func init() {
	LocationTimes = 28800 //东八区
	LoadLocation()
}

//设置时区
func LoadLocation(location ...string) (err error) {
	if len(location) < 1 {
		cstZone, err = time.LoadLocation("Local")
	} else {
		cstZone, err = time.LoadLocation(location[0])
	}
	return
}

//just remind
//当前时间戳
func Now() int64 {
	return time.Now().Unix()
}

//当前时间格式输出
func NowFormat(format string) string {
	return time.Now().Format(format)
}

//今天星期几
func NowWeekday() string {
	return time.Now().Weekday().String()
}

//day

//Get the timestamp of the midnight in cn
//查询当天零点时间戳，注意时区,减去八个小时
func GetCNTodayStartTs() (int64, error) {
	t, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		return 0, err
	}
	return t.Unix() - LocationTimes, nil
}

//Get the timestamp of the midnight ， pay attention to the time zone
//查询当天零点时间戳，注意时区,减去八个小时
func GetTodayStartTs() (int64, error) {
	t, err := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"),time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func GetTomorrowStartTs() (int64, error) {
	t, err := time.ParseInLocation("2006-01-02", time.Now().Add(24*time.Hour).Format("2006-01-02"),time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

//week
//查询本周周一零点时间
func GetNowMonday() time.Time {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	monday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return monday
}
