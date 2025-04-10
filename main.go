package main

import (
	"log"
	"study-spider-audiobook/config"
	"study-spider-audiobook/logger"

	"github.com/spf13/viper"
)

// 初始化
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
	viper.Set("network.ximalaya_Ip", "www.baidu.com")
	if err := config.WriteConfig2Blank(".", "config", "yaml"); err != nil {
		// if err := config.WriteConfig4Blank(cfg); err != nil {
		log.Fatalln("写入配置文件失败,err: ", err)
	}
}
func main() {

}
