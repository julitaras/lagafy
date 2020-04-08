package util

import "api-dashboard/pkg/setting"

//Setup inizialize the
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
