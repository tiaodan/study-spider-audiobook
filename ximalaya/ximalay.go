package ximalaya

import (
	"log"
	models "study-spider-audiobook/models/ximalaya"

	"github.com/levigross/grequests"
)

// 爬取喜马拉雅【广播剧】, 类似接口的方式。
/*
思路:
1. 人工用浏览器F12找到某一个类别的请求（如：文学名著）
2. 把该链接url, 作为方法参数

步骤:
1. 请求url, 获取json数据
*/
func SpiderGuangbojuByInterface(url string) {
	// url := "https://www.ximalaya.com/revision/category/v2/albums?pageNum=1&pageSize=56&sort=1&categoryId=15&metadataValues=%E6%96%87%E5%AD%A6%E5%90%8D%E8%91%97"

	// 带请求头写法
	// resp, err := grequests.Get(url, &grequests.RequestOptions{
	// 	Headers: map[string]string{
	// 		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	// 	},
	// })

	// 不带请求头写法
	resp, err := grequests.Get(url, nil)

	if err != nil {
		panic(err)
	}

	// 检查状态码
	if resp.StatusCode != 200 {
		log.Println("请求失败，状态码：", resp.StatusCode)
		panic("请求失败，状态码：")
	}

	var result models.XimalayaResponseData
	if err := resp.JSON(&result); err != nil {
		panic(err)
	}

	// 输出返回信息
	log.Println("返回信息: 响应码: ", result.Ret)
	log.Println("返回信息: 专辑总数: ", result.Data.Total)
	log.Println("返回信息: 当前页码: ", result.Data.PageNum)
	log.Println("返回信息: 每页个数: ", result.Data.PageSize)

	// 输出某个有声书信息 arr[0]
	album0 := result.Data.Albums[0]
	log.Println("arr[0]有声书信息: ID: ", album0.AlbumID)
	log.Println("arr[0]有声书信息: 播放量: ", album0.AlbumPlayCount)
	log.Println("arr[0]有声书信息: 音轨数量, 就是有多少集: ", album0.AlbumTrackCount)
	log.Println("arr[0]有声书信息: 封面路径: ", album0.AlbumCoverPath)
	log.Println("arr[0]有声书信息: 标题: ", album0.AlbumTitle)
	log.Println("arr[0]有声书信息: 单个有声书链接: ", album0.AlbumUrl)
	log.Println("arr[0]有声书信息: 有声书简介: ", album0.Intro)

	log.Println("arr[0]作者信息: 作者: ", album0.AlbumUserNickName)
	log.Println("arr[0]作者信息: 主持人id,: ", album0.AnchorId)
	log.Println("arr[0]作者信息: 作者分数: ", album0.AnchorGrade)
	log.Println("arr[0]作者信息: 作者主页链接: ", album0.AnchorUrl)
	log.Println("arr[0]作者信息: mvp分数: ", album0.MvpGrade)

	log.Println("arr[0]其它信息: 是否删除标志: ", album0.IsDeleted)
	log.Println("arr[0]其它信息: 是否已支付: ", album0.IsPaid)
	log.Println("arr[0]其它信息: 是否完结: ", album0.IsFinished)
	log.Println("arr[0]其它信息: logo类型: ", album0.LogoType)

	// 输出专辑信息
	// for _, album := range result.Data.Albums {
	// 	log.Printf("标题: %s", album.AlbumTitle)
	// 	log.Printf("封面链接: %s", album.AlbumCoverPath)
	// 	log.Printf("播放量: %d", album.PlayCount)
	// 	log.Println("-------------------")
	// }
}
