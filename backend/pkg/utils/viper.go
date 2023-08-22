package utils

import "flag"

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/8/22 9:19
//

// InitViper 初始化配置文件 优先级: 命令行 -> 默认值
func InitViper() {
	// 根据命令号读取配置文件路径
	var filePath string
	flag.StringVar(&filePath, "c", "", "input config file .")
	flag.Parse()

}
