// @User CPR
package utils

import (
	"VideoWeb/config"
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// 优先级：命令行参数 > 环境变量 > 配置文件
func InitViper() {
	// 根据命令行读取配置文件路径
	var configPath string
	flag.StringVar(&configPath, "c", "", "choose config file.")
	// 参数传入configPath 命令行 通过-c传入 默认为空
	flag.Parse()
	if configPath != "" { // 命令行读取到参数
		log.Printf("命令行读取参数, 配置文件路径为: %s\n", configPath)
	} else { // 命令行未读取到参数
		log.Printf("命令行参数为空, 默认加载: config/config.toml\n")
		configPath = "config/config.toml"
	}

	// 读取固定路径的配置文件
	// viper是一个配置管理的解决方案，它能够从 json，toml，ini，yaml，hcl，env 等多种格式文件中，
	// 读取配置内容，它还能从一些远程配置中心读取配置文件，如consul，etcd等；它还能够监听文件的内容变化。
	v := viper.New()
	v.SetConfigFile(configPath)
	v.AutomaticEnv()                                   // 允许使用环境变量
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // SERVER_APPMODE => SERVER.APPMODE

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Panic("配置文件读取失败, err: ", err)
	}

	// 加载配置文件内容到结构体对象
	if err := v.Unmarshal(&config.Cfg); err != nil {
		log.Panic("配置文件解析失败, err: ", err)
	}
	// todo：配置文件热重载，使用场景是什么？
	v.WatchConfig() // 监听配置文件变化
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("检测到配置文件变化，重新加载配置文件")
		if err := v.Unmarshal(&config.Cfg); err != nil {
			log.Panic("配置文件解析失败, err: ", err)
		}
	})
	log.Println("配置文件加载成功")
}
