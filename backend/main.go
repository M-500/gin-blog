package main

import (
	"backend/models"
	"backend/pkg/config"
	"backend/pkg/utils"
	"backend/pkg/utils/sqls"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func main() {
	// 1. 初始化全局变量
	fmt.Println("fuck you")
}

func init() {
	// 1. 初始化配置
	utils.InitViper()

	utils.InitLogger()

	// 3. 初始化MySQL的连接
	gormConf := &gorm.Config{
		Logger: logger.New(logrus.StandardLogger(), logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
		}),
	}

	if err := sqls.Open(config.Cfg.Mysql, gormConf, models.Models...); err != nil {
		logrus.Error(err)
	}

	// 4. 初始化Redis的连接
}
