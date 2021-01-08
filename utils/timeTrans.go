/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2021-01-05 20:29:43
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-07 09:33:24
 */
package utils

import (
	"time"
)

var timeTemplate1 = "2006-01-02"          //常规类型1
var timeTemplate2 = "2006-01-02 15:04:05" //常规类型2

func String2timeStamp(origin string) int64 {
	tmp, _ := time.ParseInLocation(timeTemplate1, origin, time.Local)
	return tmp.Unix()
}

func TimeStamp2string(timestamp int64) string {
	str := time.Unix(timestamp, 0).Format(timeTemplate1)
	return str
}

//detail
func String2timeStamp2(origin string) int64 {
	tmp, _ := time.Parse(timeTemplate2, origin)
	return tmp.Unix()
}

//detail
func TimeStamp2string2(timestamp int64) string {
	str := time.Unix(timestamp, 0).Format(timeTemplate2)
	return str
}
