package main

import (
	"backend/models"
	"backend/pkg/config"
	"backend/pkg/utils"
	"backend/pkg/utils/sqls"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var g errgroup.Group

func main() {
	// 1. 初始化全局变量
	fmt.Println("fuck you")

	// 前台接口服务
	g.Go(func() error {
		return nil
	})

	// 后台接口服务
	g.Go(func() error {
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatal()
	}
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
