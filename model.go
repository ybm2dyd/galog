package galog

import (
	"fmt"
	"path"
)

// Playerlogin 玩家登录
type Playerlogin struct {
	ZoneID        int    // 游戏区编号
	EventTime     string // 事件时间(服务器本地时间)如 2020-01-01 12:23:45
	Timestamp     int64  // 毫秒级时间戳
	GameID        int    // 游戏id
	ChannelID     int    // 渠道id
	AccountID     string // 账号id 形如 1-1233
	PlatID        int    // 客户端平台 ios:2/android:1
	CharID        string // 角色id
	CharName      string // 角色名字 UTF-8编码
	DeviceID      string // 设备id(iOS上是idfa android上是OAIDimei_md5等)
	ClientVersion string // 客户端版本
	ClientIP      string // 客户端IP地址
	Level         int    // 角色等级
	VipLevel      int    // 角色vip等级
	Career        int    // 角色职业
	Regtime       string // 角色创建时间，服务器本地时间
	CharType      int    // 角色类型： 正常:1 测试:2 gm/福利:3 机器人/AI:4 其他:5
	OS            string // 软件版本 操作系统版本
	PhoneModel    string // 硬件机型 品牌-型号
	Operator      string // 运营商 mobile/telecom/unicom
	Network       string // 网络 WIFI/2G/3G/4G/5G
	CPU           string // cpu类型 cpu类型_频率_核数等
	Memory        string // 内存，单位M
	PageID        int    // 角色所在场景id
	PageName      string // 角色所在场景名称
	KvGroup       string // 扩展字段，一个或多个KeyValue的JSON字符串
}

// Playerlogout 玩家登出
type Playerlogout struct {
	ZoneID        int    // 游戏区编号
	EventTime     string // 事件时间(服务器本地时间)如 2020-01-01 12:23:45
	Timestamp     int64  // 毫秒级时间戳
	GameID        int    // 游戏id
	ChannelID     int    // 渠道id
	AccountID     string // 账号id 形如 1-1233
	PlatID        int    // 客户端平台 ios:2/android:1
	CharID        string // 角色id
	CharName      string // 角色名字 UTF-8编码
	DeviceID      string // 设备id(iOS上是idfa android上是OAIDimei_md5等)
	ClientVersion string // 客户端版本
	ClientIP      string // 客户端IP地址
	Level         int    // 角色等级
	VipLevel      int    // 角色vip等级
	Career        int    // 角色职业
	OnlineTime    int    // 角色本次在线时间(秒)
	CharType      int    // 角色类型： 正常:1 测试:2 gm/福利:3 机器人/AI:4 其他:5
	Reason        int    // 下线原因： 正常:1 被gm踢:2 被自己踢:3 异常下线:4 其他;5
	OS            string // 软件版本 操作系统版本
	PhoneModel    string // 硬件机型 品牌-型号
	Operator      string // 运营商 mobile/telecom/unicom
	Network       string // 网络 WIFI/2G/3G/4G/5G
	CPU           string // cpu类型 cpu类型_频率_核数等
	Memory        string // 内存，单位M
	PageID        int    // 角色所在场景id
	PageName      string // 角色所在场景名称
	KvGroup       string // 扩展字段，一个或多个KeyValue的JSON字符串
}

var playerloginLogger, playerlogoutLogger *Logger

// Init sdk init
func Init(logpath string, rollingtime Rollingtime, rollinginterval int) error {
	var err error
	playerloginLogger, err = getLogger(logpath, "playerlogin", rollingtime, rollinginterval)
	if err != nil {
		return err
	}

	playerlogoutLogger, err = getLogger(logpath, "playerlogout", rollingtime, rollinginterval)
	if err != nil {
		return err
	}
	return nil
}

func getLogger(logpath string, ident string, rollingtime Rollingtime, rollinginterval int) (*Logger, error) {
	logger := new(Logger)
	formatter := TextFormatter{
		DisableFormat: true,
	}
	logger.SetFormatter(&formatter)
	logger.SetNoLock()
	output, err := NewTimeRotatingFileHandler(path.Join(logpath, ident), rollingtime, rollinginterval)
	if err != nil {
		fmt.Println("[ERROR]", err.Error())
		return nil, err
	}
	logger.SetOutput(output)
	logger.SetLevel(InfoLevel)
	return logger, nil
}

// Log Playerlogin 写日志
func (p Playerlogin) Log() error {
	return playerloginLogger.Infof("Playerlogin - %d|%s|%d|%d|%d|%s|%d|%s|%s|%s|%s|%s|%d|%d|%d|%s|%d|%s|%s|%s|%s|%s|%s|%d|%s|%s\n",
		p.ZoneID,
		p.EventTime,
		p.Timestamp,
		p.GameID,
		p.ChannelID,
		p.AccountID,
		p.PlatID,
		p.CharID,
		p.CharName,
		p.DeviceID,
		p.ClientVersion,
		p.ClientIP,
		p.Level,
		p.VipLevel,
		p.Career,
		p.Regtime,
		p.CharType,
		p.OS,
		p.PhoneModel,
		p.Operator,
		p.Network,
		p.CPU,
		p.Memory,
		p.PageID,
		p.PageName,
		p.KvGroup)
}

// Log Playerlogout 写日志
func (p Playerlogout) Log() error {
	return playerlogoutLogger.Infof("Playerlogout - %d|%s|%d|%d|%d|%s|%d|%s|%s|%s|%s|%s|%d|%d|%d|%d|%d|%d|%s|%s|%s|%s|%s|%s|%d|%s|%s\n",
		p.ZoneID,
		p.EventTime,
		p.Timestamp,
		p.GameID,
		p.ChannelID,
		p.AccountID,
		p.PlatID,
		p.CharID,
		p.CharName,
		p.DeviceID,
		p.ClientVersion,
		p.ClientIP,
		p.Level,
		p.VipLevel,
		p.Career,
		p.OnlineTime,
		p.CharType,
		p.Reason,
		p.OS,
		p.PhoneModel,
		p.Operator,
		p.Network,
		p.CPU,
		p.Memory,
		p.PageID,
		p.PageName,
		p.KvGroup)
}

func Clean() {
	playerloginLogger.Out.Close()
	playerlogoutLogger.Out.Close()
}
