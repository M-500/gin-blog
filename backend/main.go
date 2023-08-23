package main

import (
	"backend/app/svc"
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
)

var g errgroup.Group

func main() {
	// 1. 初始化全局变量
	var (
		configPath string
		port       int
		initDb     bool
	)
	// 1. 读取配置文件路经
	flag.StringVar(&configPath, "cfg", "etc/dev.yml", "配置文件路径")
	// 2. 读取端口信息
	flag.IntVar(&port, "port", 8372, "端口号")
	// 3. 读取是否初始化数据库信息
	flag.BoolVar(&initDb, "setup", false, "初始化项目相关数据")
	flag.Parse()

	serSvc := svc.NewSerContext(configPath)
	err := serSvc.BeforeStart()
	if err != nil {
		return
	}

	// 前台接口服务
	g.Go(func() error {
		return nil
	})
	addr := fmt.Sprintf("%s:%d", serSvc.Config.Server.Host, port)
	// 后台接口服务
	g.Go(func() error {
		err = serSvc.Server.Run(addr)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatal()
	}
}

func init() {
	// 1. 初始化配置
	//utils.InitLogger()
	//
	//// 3. 初始化MySQL的连接
	//gormConf := &gorm.Config{
	//	Logger: logger.New(logrus.StandardLogger(), logger.Config{
	//		SlowThreshold:             time.Second,
	//		Colorful:                  true,
	//		LogLevel:                  logger.Warn,
	//		IgnoreRecordNotFoundError: true,
	//	}),
	//}
	//
	//if err := sqls.Open(config.Cfg.Mysql, gormConf, models.Models...); err != nil {
	//	logrus.Error(err)
	//}

	// 4. 初始化Redis的连接
}
