package main

import (
	"study-spider-audiobook/config"
	"study-spider-audiobook/logger"
	"study-spider-audiobook/ximalaya"
)

// 初始化, 默认main会自动调用本方法
func init() {
	// 获取配置实例（首次调用时触发初始化）
	cfg := config.GetConfig(".", "config", "yaml")

	// 读取配置文件，并设置为日志级别, 默认info
	switch cfg.Log.Level {
	case "debug":
		logger.SetLogLevel(logger.LevelDebug)
	case "info":
		logger.SetLogLevel(logger.LevelInfo)
	case "warn":
		logger.SetLogLevel(logger.LevelWarn)
	case "error":
		logger.SetLogLevel(logger.LevelError)
	default:
		logger.SetLogLevel(logger.LevelInfo)
	}

	// 打印配置
	logger.Debug("network.ximalayaIIp_ip: %s", cfg.Network.XimalayaIIp)
}

/*
思路:
 1. 读取配置文件， (如果配置文件不填, 自动会有默认值)
 2. 设置日志级别, 默认info
 3. 统一调用错误打印, 封装函数
 4. 爬取页面数据, 尽量去重
 5. 插入数据库
*/
func main() {
	// 1. 读取配置文件， (如果配置文件不填, 自动会有默认值)
	// 2. 设置日志级别, 默认info
	// 3. 统一调用错误打印, 封装函数
	// 4. 爬取页面数据, 尽量去重
	url := "https://www.ximalaya.com/revision/category/v2/albums?pageNum=1&pageSize=56&sort=1&categoryId=15&metadataValues=%E6%96%87%E5%AD%A6%E5%90%8D%E8%91%97"
	ximalaya.SpiderGuangbojuByInterface(url)
}
