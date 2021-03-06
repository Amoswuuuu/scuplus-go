package model

const (
	// LostFindReady 准备发布
	LostFindReady = 0
	// LostFindShow 已发布
	LostFindShow = 1

	LostFindCard = "一卡通招领"
)

// LostFind 失物招领
type LostFind struct {
	Model
	UserID   uint   `json:"user_id"` // 用户id
	Title    string `json:"title"`   // 标题
	Nickname string `json:"nickname"`
	Pictures string `json:"pictures"`  // 截图链接
	Info     string `json:"info"`      // 信息
	Address  string `json:"address"`   // 地点
	Contact  string `json:"contact"`   // 联系方式
	CardInfo string `json:"card_info"` // 一卡通信息,json字符串仅一卡通有
	Category string `json:"category"`  // 分类: 一卡通,其他,遗失
	Status   int    `json:"status"`    // 0: 准备中, 1: 已发布
}
