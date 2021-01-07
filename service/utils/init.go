/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-12-08 19:50:35
 * @LastEditors: Seven
 * @LastEditTime: 2020-12-08 19:53:32
 */
package utils

func init() {
	LoadLocalConfig()
	// LoadConsulConfig()
	// ticker := time.NewTicker(time.Duration(LocalConf.ConfigTTL) * time.Second)
	// go func() {
	// 	for {
	// 		<-ticker.C
	// 		LoadConsulConfig()
	// 	}
	// }()

	LoadLog()
}
