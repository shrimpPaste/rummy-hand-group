package setting

import (
	"time"
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	//DefaultPageSize: 10
	//MaxPageSize: 100
	//LogSavePath: strong/logs
	//LogFileName: app
	//LogFileExt: .log

	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	Main    MySQLSettingS `mapstructure:"main"`
	Ch      MySQLSettingS `mapstructure:"ch"`
	Slave   MySQLSettingS `mapstructure:"slave"`
	Conf    MySQLSettingS `mapstructure:"conf"`
	FlowM   MySQLSettingS `mapstructure:"flow_m"`
	FlowS   MySQLSettingS `mapstructure:"flow_s"`
	BetM    MySQLSettingS `mapstructure:"bet_m"`
	RobotM  MySQLSettingS `mapstructure:"robot_m"`
	BetS    MySQLSettingS `mapstructure:"bet_s"`
	RobotS  MySQLSettingS `mapstructure:"robot_s"`
	DcM     MySQLSettingS `mapstructure:"dc_m"`
	DCS     MySQLSettingS `mapstructure:"dc_s"`
	ExeM    MySQLSettingS `mapstructure:"exe_m"`
	ExeS    MySQLSettingS `mapstructure:"exe_s"`
	PlayerM MySQLSettingS `mapstructure:"player_m"`
	PlayerS MySQLSettingS `mapstructure:"player_s"`
	BuryM   MySQLSettingS `mapstructure:"bury_m"`
	BuryS   MySQLSettingS `mapstructure:"bury_s"`
	//DBType       string
	//UserName     string
	//Password     string
	//Host         string
	//DBName       string
	//UserDbName   string // 存储用户的数据库
	//TablePrefix  string
	//Charset      string
	//ParseTime    bool
	//MaxIdleConns int
	//MaxOpenConns int
}

type RedisS struct {
	Sess []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	} `mapstructure:"sess"`
	DataAi []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	} `mapstructure:"data_ai"`
	Data []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	}
	Plays []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	}
	Cache []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	}
	Temp []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	}
	Ttl []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	}
	Bet []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	}
	Play []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	}
	RotCache []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	} `mapstructure:"rot_cache"`
	LogAgent []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	} `mapstructure:"log_agent"`
	Rank []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	}
	User []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	}
	Cross []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	}
	Dealer []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	}
	Tmp []struct {
		Key           string
		RedisSettingS `mapstructure:",squash"`
	}
}

type RedisSettingS struct {
	Host string `yaml:"host" mapstructure:"host"`
	Port int    `yaml:"port" mapstructure:"port"`
	Pass string `yaml:"pass" mapstructure:"pass"`
}

type MySQLSettingS struct {
	Host     string `yaml:"host" mapstructure:"host"`
	UserName string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
	DBName   string `yaml:"dbname" mapstructure:"dbname"`
	Charset  string `yaml:"charset" mapstructure:"charset"`
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
