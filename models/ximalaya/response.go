// 喜马拉雅 接口响应数据模型, 存数据用的
package models

type XimalayaResponseData struct {
	Ret  int `json:"ret"` // 响应状态码
	Data struct {
		Total    int             `json:"total"`    // 专辑总数, 某一类总数, 如文学名著
		PageNum  int             `json:"pageNum"`  // 当前页码
		PageSize int             `json:"pageSize"` // 每页个数
		Albums   []XimalayaAlbum `json:"albums"`   // 专辑,喜马拉雅叫这个名, 自己起就叫AudioBook,好记
	} `json:"data"`
}
