package global

import (
	"trading-system/pkg/setting"
	"trading-system/pkg/logger"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger			*logger.Logger
)
