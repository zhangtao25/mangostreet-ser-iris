package app

import (
	"mangostreet-ser-iris/common"
)

func StartOn() {
	if !common.IsProd() {
		return
	}

	// 开启定时任务
	startSchedule()
}
