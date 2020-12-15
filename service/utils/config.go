/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-12-08 19:21:22
 * @LastEditors: Seven
 * @LastEditTime: 2020-12-08 19:52:27
 */
package utils

import (
	"os"
	"path"
	"runtime"

	"github.com/micro/go-micro/config"
)

type Config struct {
	Deploy    string            `json:"deploy"`
	EctdAddr  string            `json:"etcdAddr"`
	ConfigTTL int               `json:"config_ttl"`
	Hosts     map[string]string `json:"hosts"`
}

var LocalConf Config

// LoadLocalConfig load config from local
func LoadLocalConfig() {
	_, filename, _, ok := runtime.Caller(0) //file name会返回调用caller的地方，也就是util
	LogPanic(ok, "No caller information")
	//获取环境变量
	p := os.Getenv("JUB_CONFIGPATH") //没有定义环境变量，就在项目根目录下找
	if p != "" {
		LogPanic(config.LoadFile(p))
	} else {
		LogPanic(config.LoadFile(path.Dir(filename) + "/../config.json"))
	}
	LogPanic(config.Scan(&LocalConf))
	Info("Local config loaded")
}
