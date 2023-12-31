package utils

import (
	"backend/app/config"
	"backend/pkg/utils/strs"
	"flag"
	"github.com/spf13/viper"
	"log"
	"strings"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/8/22 9:19
//

// InitViper 初始化配置文件 优先级: 命令行 -> 默认值
func InitViper(path string) (*config.Config, error) {
	var cfg *config.Config
	// 根据命令号读取配置文件路径
	var filePath string
	flag.StringVar(&filePath, "c", "", "input config file .")
	flag.Parse()

	if !strs.IsBlank(filePath) {
		log.Printf("命令行读取参数, 配置文件路径为: %s\n", filePath)
	} else {
		log.Println("命令行参数为空, 默认加载: ", path)
		filePath = path
	}
	// 目前读取固定固定路径的配置文件
	v := viper.New()
	v.SetConfigFile(filePath)
	v.AutomaticEnv()                                   // 允许使用环境变量
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // SERVER_APPMODE => SERVER.APPMODE
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Panic("配置文件读取失败: ", err)
	}

	// 加载配置文件内容到结构体对象
	if err := v.Unmarshal(&cfg); err != nil {
		log.Panic("配置文件内容加载失败: ", err)
		return nil, err
	}
	// TODO: 配置文件热重载, 使用场景是什么?
	// v.WatchConfig()
	// v.OnConfigChange(func(e fsnotify.Event) {
	// 	log.Println("检测到配置文件内容修改")
	// 	if err := v.Unmarshal(&config.Cfg); err != nil {
	// 		log.Panic("配置文件内容加载失败: ", err)
	// 	}
	// })
	log.Println("配置文件内容加载成功")
	return cfg, nil
}
