package global

import (
	"rummy-logic-v3/pkg/logger"
	"rummy-logic-v3/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	RedisSetting    *setting.RedisS
	Logger          *logger.Logger
)
