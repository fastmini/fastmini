/*
 * @Author: Ali2vu <751815097@qq.com>
 * @Date: 2021-08-10 00:26:55
 * @LastEditors: Ali2vu
 * @LastEditTime: 2021-08-10 01:20:23
 */
package config

import "os"

func Config(key string, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
