/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2021-01-06 11:49:31
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-06 11:55:29
 */
package utils

func StringToInt32Arr(s string, arr []int32) {
	src := []int32(s)
	j := 0
	for i, v := range src {
		if i >= len(arr) {
			break
		}
		if v == ',' || v == '[' || v == ']' {
			continue
		}
		arr[j] = v
		j++
	}
}
