package galog

import (
	"fmt"
	"path"
)

// Playerlogin 玩家登录
type Playerlogin struct {
	ZoneID        int    // 游戏区ID 游戏服务器编号
	EventTime     string // 游戏事件的时间, 服务器本地时间，格式 YYYY-MM-DD HH:MM:SS
	Timestamp     int64  // 时间戳，到毫秒，例如：1585903230000
	GameID        int    // 游戏id,统一分配，从平台部申请，如果发行多个地区，建议按地区申请
	ChannelID     int    // 渠道ID，来源渠道配置表
	AccountID     string // 账号ID，按照平台统一规则，与客户端sdk的openid一致，比如1-33333
	PlatID        int    // 客户端平台 ios 2/android 1
	CharID        string // 角色id
	CharName      string // 角色名字 UTF-8编码
	DeviceID      string // 设备ID（ios的idfa，android的imei或mac等）
	ClientVersion string // 客户端版本
	ClientIP      string // 客户端IP地址
	Level         int    // 角色触发当前事件时的等级
	VipLevel      int    // 角色vip等级
	Career        int    // 角色职业
	Regtime       string // 角色注册时间，在游戏服生成角色ID时的时间，服务器本地时间，格式 YYYY-MM-DD HH:MM:SS
	CharType      int    // 角色类型，标识角色的分类属性，1:正常 2:测试 3:gm/福利 4:AI玩家  5:其他
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

var playerloginLogger *Logger

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

// Playerlogout 玩家登出
type Playerlogout struct {
	ZoneID        int    // 游戏区编号
	EventTime     string // 游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	Timestamp     int64  // 时间戳，到毫秒，例如：1585903230000
	GameID        int    // 游戏id
	ChannelID     int    // 渠道ID，来源渠道配置表
	AccountID     string // 账号ID，按照平台统一规则，与客户端sdk的openid一致，比如1-33333
	PlatID        int    // 客户端平台 ios:2/android:1
	CharID        string // 角色id
	CharName      string // 角色名字 UTF-8编码
	DeviceID      string // 设备ID（ios的idfa，android的imei或mac等），更多的可以放在自定义字段中
	ClientVersion string // 客户端版本
	ClientIP      string // 客户端IP地址
	Level         int    // 角色触发当前事件时的等级
	VipLevel      int    // 角色vip等级
	Career        int    // 角色职业
	OnlineTime    int    // 本次累计在线时间(秒)
	CharType      int    // 角色类型，标识角色的分类属性，1:正常 2:测试 3:gm/福利 4:AI玩家  5:其他
	Reason        int    // 1: 正常登出 2: gm踢下线 3: 挤下线（同账号多处登录）4:异常退出 5: 其他
	OS            string // 软件版本 操作系统版本
	PhoneModel    string // 硬件机型 品牌;型号
	Operator      string // 运营商 mobile/telecom/unicom
	Network       string // 网络 WIFI/2G/3G/4G/5G
	CPU           string // cpu类型 cpu类型_频率_核数
	Memory        string // 内存，单位M
	PageID        int    // 角色所在场景id
	PageName      string // 角色所在场景名称
	KvGroup       string // 扩展字段，一个或多个KeyValue的JSON字符串
}

var playerlogoutLogger *Logger

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

// Roundflow 对战流水
type Roundflow struct {
	ZoneID     int    // 游戏区编号
	EventTime  string // 游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	Timestamp  int64  // 时间戳，到毫秒，例如：1585903230000
	GameID     int    // 游戏id
	AccountID  string // 账号ID，按照平台统一规则，与客户端sdk的openid一致，比如1-33333
	PlatID     int    // 客户端平台 ios:2/android:1
	CharID     string // 角色id
	CharName   string // 角色名字 UTF-8编码
	Level      int    // 挑战的等级
	BattleType string // 玩法类型中文名 如：主线关卡
	RoundID    int    // 本局唯一 id
	BattleID   string // 关卡 ID 或副本 ID
	BattleName string // 玩法中文名
	FightPoint int    // 本局结束时战力
	RoundScore int    // 本局分数 得分（无为空）
	RoundTime  int    // 对局时长(秒) 时长
	Result     int    // 单局结果 1胜利，2失败
	rank       int    // 排名 排名（无为空）
	KvGroup    string // 自定义KeyValue字段组 json 格式记录挑战宠物、技能以及队伍信息{"pet":宠物ID ,"heroskill":[{"skillcid":3017104,"skilllv":52,"skillquality":0},{"skillcid":3017103,"skilllv":52,"skillquality":0}],"id":"team1234","members":[{"teammate"：队友账号ID，"career"：队友职业名称}]}
}

var roundflowLogger *Logger

// Log Roundflow 写日志
func (p Roundflow) Log() error {
	return roundflowLogger.Infof("Roundflow - %d|%s|%d|%d|%s|%d|%s|%s|%d|%s|%d|%s|%s|%d|%d|%d|%d|%d|%s\n",
		p.ZoneID,
		p.EventTime,
		p.Timestamp,
		p.GameID,
		p.AccountID,
		p.PlatID,
		p.CharID,
		p.CharName,
		p.Level,
		p.BattleType,
		p.RoundID,
		p.BattleID,
		p.BattleName,
		p.FightPoint,
		p.RoundScore,
		p.RoundTime,
		p.Result,
		p.rank,
		p.KvGroup)
}

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

	roundflowLogger, err = getLogger(logpath, "roundflow", rollingtime, rollinginterval)
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

// Clean loggers clean
func Clean() {

	playerloginLogger.Out.Close()
	playerlogoutLogger.Out.Close()
	roundflowLogger.Out.Close()
}
