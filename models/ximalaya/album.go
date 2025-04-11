// // 喜马拉雅 接口响应数据模型, 存数据用的
package models

// 专辑中,某张唱片
type XimalayaAlbum struct {
	// 有声书相关
	AlbumID         int    `json:"albumId"`         // 有声书id
	AlbumPlayCount  int    `json:"albumPlayCount"`  // 播放量
	AlbumTrackCount int    `json:"albumTrackCount"` // 音轨数量, 就是有多少集
	AlbumCoverPath  string `json:"albumCoverPath"`  // 封面路径
	AlbumTitle      string `json:"albumTitle"`      // 标题, 显示: 西游记 | 免费广播剧
	AlbumUrl        string `json:"albumUrl"`        // 单个有声书链接
	Intro           string `json:"intro"`           // 有声书简介

	// 作者相关
	AlbumUserNickName string `json:"albumUserNickName"` // 作者
	AnchorId          int    `json:"anchorId"`          // 作者id, 就是主持人id,Anchor 主持人的意思
	AnchorGrade       int    `json:"anchorGrade"`       // 作者分数，好像1-15，15最高
	AnchorUrl         string `json:"anchorUrl"`         // 作者主页链接
	MvpGrade          int    `json:"mvpGrade"`          // mvp分数，好像1-?

	// 其它待分类
	IsDeleted  bool `json:"isDeleted"`  // 是否删除标志
	IsPaid     bool `json:"isPaid"`     // 是否已支付
	IsFinished int  `json:"isFinished"` // 是否完结
	VipType    int  `json:"vipType"`    // vip类型？不知道干嘛用的
	LogoType   int  `json:"logoType"`   // logo类型？不知道干嘛的
	// 其他字段按需添加...
}
