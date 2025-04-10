package main

import (
	"log"
	"study-spider-audiobook/config"
	"study-spider-audiobook/logger"
)

func main() {
	// 获取配置实例（首次调用时触发初始化）
	cfg := config.GetConfig(".", "config", "yaml")

	// 使用配置
	log.Println("network.ximalayaIIp_ip: ", cfg.Network.XimalayaIIp)

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

	logger.Debug("这是一条调试信息")
	logger.Info("这是一条普通信息")
	logger.Warn("这是一条警告信息")
	logger.Error("这是一条错误信息")

	// 修改配置
	cfg.Network.XimalayaIIp = "8.8.8.8"
	// 写入配置文件
	// if err := config.WriteConfig(cfg); err != nil {
	// 	log.Fatalln("写入配置文件失败,err: ", err)
	// }
	if err := config.WriteConfig2Blank(); err != nil {
		log.Fatalln("写入配置文件失败,err: ", err)
	}
}
