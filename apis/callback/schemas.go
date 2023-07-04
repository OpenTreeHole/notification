package callback

type MipushCallbackData struct {
	// 开发者自定义参数
	Param string `json:"param"`

	// 消息状态类型。
	// 1-送达；2-点击；16-无法找到目标设备，32-客户端调用了disablePush接口禁用Push，
	// 64-目标设备不符合过滤条件，128-当日推送总量超限或单设备接收超限
	Type int `json:"type"`

	// targets: 一批alias、regId或useraccount列表, 逗号分隔
	Targets string `json:"targets"`

	// 发送消息时设置的jobkey值
	Jobkey string `json:"jobkey"`

	// 消息送达时通知栏的状态
	// “Enable”为用户允许此app展示通知栏消息， “Disable”为通知栏消息已关闭。 “Unknown”通知栏状态未知。
	BarStatus string `json:"barStatus"`

	// 消息送到设备的时间
	Timestamp int `json:"timestamp"`

	// 表示该发送目标无法正常送达，但建议可替换成另外一个新的目标标识以满足该设备的正常推送。
	// 其值为键值对，key表示该设备原始发送目标标识，value表示建议替换成的新的目标标识。
	ReplaceTarget map[string]string `json:"replaceTarget"`

	// errorCode
	// type为16时返回的无效目标的子类。
	// errorCode:1表示无效regid
	// errorCode:2表示无效alias
	// errorCode:3表示无效useraccount
	ErrorCode int `json:"errorCode"`
}
