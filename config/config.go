package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// 配置文件 结构体
type Config struct {
	Network struct {
		LocalIp    string `mapstructure:"local_ip"`
		LocalPort  string `mapstructure:"local_port"`
		RemoteIp   string `mapstructure:"remote_ip"`
		RemotePort string `mapstructure:"remote_port"`
	}
}

var (
	cfg  *Config   // 全局变量
	once sync.Once //保证单例初始化
)

// GetConfig 获取配置实例（单例）
/*
思路:
   1. 初始化Viper
   2. 设置默认值, 防止用户没配置, 读取到空值
   3. 读取配置文件
   4. 将配置文件解析到结构体
   5. 返回配置指针
*/
func GetConfig() *Config {
	once.Do(func() {
		// 初始化Viper
		viper.AddConfigPath(".")      //配置文件搜索路径（当前目录）
		viper.SetConfigName("config") // 配置文件名（不含扩展名）
		viper.SetConfigType("ini")    // 文件类型（yaml、json 等）

		// 设置默认值, 防止用户没配置, 读取到空值
		viper.SetDefault("network.local_ip", "127.0.0.1")
		viper.SetDefault("network.local_port", "8080")
		viper.SetDefault("network.remote_ip", "192.168.85.93")
		viper.SetDefault("network.remote_port", "3200")

		// 读取配置文件
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalln("读取配置文件失败,err: ", err)
		}

		// 将配置文件解析到结构体
		cfg = &Config{}
		if err := viper.Unmarshal(cfg); err != nil {
			log.Fatalln("解析配置文件失败,err: ", err)
		}
	})
	return cfg
}

// WriteConfig 将配置写入文件
func WriteConfig(cfg *Config) error {
	// 将结构体转换为 map[string]interface{}
	var newCfg map[string]interface{}
	if err := mapstructure.Decode(cfg, &newCfg); err != nil {
		return fmt.Errorf("结构体转 Map 失败:: %v", err)
	}

	// 将 Map 合并到 Viper
	viper.MergeConfigMap(newCfg)

	// 将Viper配置写入文件
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("写入配置文件失败: %v", err)
	}
	return nil
}
