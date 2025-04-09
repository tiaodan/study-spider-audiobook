package main

import (
	"fmt"
	"log"
	"study-spider-audiobook/config"
	"study-spider-audiobook/logger"
)

func main() {
	fmt.Print("Hello, World!")
	// 获取配置实例（首次调用时触发初始化）
	cfg := config.GetConfig(".", "config", "yaml")

	// 使用配置
	log.Println("network.local_ip: ", cfg.Network.LocalIp)
	log.Println("network.local_port: ", cfg.Network.LocalPort)
	log.Println("network.remote_ip: ", cfg.Network.RemoteIp)
	log.Println("network.remote_port: ", cfg.Network.RemotePort)

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
}
