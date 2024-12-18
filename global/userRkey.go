package global

// 用户相关的key
const (
	UserKey            = "hu:%v"
	KycCheck           = "kyc_check"
	IsOrganic          = "is_organic"
	StKey              = "st:%v"
	SessKey            = "sess:%v"
	TelNoKey           = "tel_no:%v"
	TelPerKey          = "tel_per"
	TelCode            = "signup_code:%v:%v"
	ConcurrenKey       = "concurren:%v:%v"
	PwdTokenKey        = "pwdtoken:%v"
	FbToken            = "fb_token:%v"
	CouponUid          = "coupon_uid:%v"
	CouponTrId         = "coupon_tr_id:%v"
	PTag               = "pTag"
	Age                = "age:%v:%v"
	ChAge              = "chAge:%v:%v:%v"
	Pay                = "pay:%v:%v:%v"
	LastRmb            = "LastRmb:%v"
	Limit              = "limit:%v:%v:%v"
	Txweb              = "txweb:%v"
	TxCfg              = "tx_cfg"
	CommIncr           = "comm_incr:%v:%v:%v"
	DayRebate          = "dayRebate:%v"
	GameData           = "vGameData:%v:%v"
	UGameData          = "uGameData:%v:%v:%v"
	ChKey              = "chKey:%v:%v"
	ChUKey             = "chUKey:%v:%v:%v"
	PlayerPlay         = "player_play:%v"
	ReddotUidKey       = "reddot:uid:%v"
	ReddotActKey       = "reddot:act:%v"
	UserInviteActivity = "invite_activity:date:%v:fromuid:%v"
	AgentDailyRebate   = "agent_Rebate:date:%v:fromuid:%v"
	GameAiList         = "ai:list:%v"
	TodayWin           = "today:win:%v:%v:%v" //taoday:win:time:vid:uid
	BetWinStat         = "bet:win:%v:%v"      // bet:win:vid:uid
	UserList           = "currency:list:%v"   // currency:list:vid
	HotUPKey           = "hotup"
	ShareUrlNew        = "share_url_new:%v"
	ShareCodeKey       = "share_code:%v"
	BetStatVKey        = "vKey:%v:%v:%v"
	BetStatUkey        = "uKey:%v:%v:%v:%v"
	UidBetKey          = "uidBetKey:%v:%v"
	UserBet            = "act_10:user_bet:%v:%v"
	UserAllBet         = "act_10:user_all_bet:%v:%v"
	UserFragments      = "act_10:user_fragments:%v:%v"
	UserFragmentsAll   = "act_10:user_fragments_all:%v:%v"
	UserIntegral       = "act_14:user_integral:%v:%v"
	UserStatKey        = "act_10:stat%v:%v:%v"
	UserByChannelKey   = "act_10:currency%v:%v:%v:%v"
	ActRankKey         = "act_10:rank:%v"
	UserSupplier       = "hu:%v:supplier"
)

//func GetCk(key string, val any) string {
//	return fmt.Sprintf("%v%v", key, val)
//}
