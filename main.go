package main

import (
	"fmt"
	"log"
	"study-spider-audiobook/config"
)

func main() {
	fmt.Print("Hello, World!")
	// 获取配置实例（首次调用时触发初始化）
	cfg := config.GetConfig()

	// 使用配置
	log.Println("network.local_ip: ", cfg.Network.LocalIp)
	log.Println("network.local_port: ", cfg.Network.LocalPort)
	log.Println("network.remote_ip: ", cfg.Network.RemoteIp)
	log.Println("network.remote_port: ", cfg.Network.RemotePort)

	// 修改配置
	cfg.Network.LocalIp = "8.8.8.8"
	// 写入配置文件
	if err := config.WriteConfig(cfg); err != nil {
		log.Fatalln("写入配置文件失败,err: ", err)
	}
}
