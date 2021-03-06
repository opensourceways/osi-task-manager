package common

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
	"reflect"
	"strings"
	"time"
)

const DATE_FORMAT = "2006-01-02 15:04:05"

func GetCurTime() string {
	return time.Now().Format(DATE_FORMAT)
}

func Catchs() {
	if err := recover(); err != nil {
		logs.Error("The program is abnormal, err: ", err)
	}
}

//TrimString Remove the \n \r \t spaces in the string
func TrimString(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	return str
}

//TrimStringNR Remove the \n \r in the string
func TrimStringNR(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	return str
}

//TimeStrToInt parse time string to unix nano
func TimeStrToInt(ts, layout string) int64 {
	if ts == "" {
		return 0
	}
	if layout == "" {
		layout = DATE_FORMAT
	}
	t, err := time.ParseInLocation(layout, ts, time.Local)
	if err != nil {
		logs.Error(err)
		return 0
	}
	return t.Unix()
}

func TimeToLocal(times, layout string) string {
	if times == "" {
		return ""
	}
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	t, err := time.ParseInLocation(layout, times, time.Local)
	if err != nil {
		logs.Error(err)
		return ""
	}
	return t.Format(DATE_FORMAT)
}

func GetAfterTime(days int) string {
	nowTime := time.Now()
	getTime := nowTime.AddDate(0, 0, days)
	resTime := getTime.Format(DATE_FORMAT) // The format of the obtained time
	logs.Info("Get the number of days：", days, ",Time ago：", resTime)
	return resTime
}

func GetCurDate() string {
	return time.Now().Format("2006-01-02")
}

func GetEnvToken(org string) string {
	org = strings.ToLower(org)
	gitToken := beego.AppConfig.String(fmt.Sprintf("repo::%s", org))
	return gitToken
}

func CreateDir(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(dir, 0777)
		}
	}
	return err
}

func SliceRemoveDup(req interface{}) (ret []interface{}) {
	if reflect.TypeOf(req).Kind() != reflect.Slice {
		return
	}
	value := reflect.ValueOf(req)
	for i := 0; i < value.Len(); i++ {
		if i > 0 && reflect.DeepEqual(value.Index(i-1).Interface(), value.Index(i).Interface()) {
			continue
		}
		ret = append(ret, value.Index(i).Interface())
	}
	return
}
