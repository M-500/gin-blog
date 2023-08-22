package svc

import (
	"backend/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/8/22 16:35
//

var svc *ServerContext

// ServerContext 保存服务相关的上下文
type ServerContext struct {
	// 全局的配置文件
	Config *config.Config
	// gin框架的引擎
	Server *gin.Engine
	// 数据库连接对象
	MysqlConn *gorm.DB
	// 翻译
	Trans ut.Translator
}

// 启动之前必须要做得事情
func (s *ServerContext) BeforeStart() error {
	// 1. 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", validators.ValidatePhone)
		// 对认证器进行翻译
		//_ = v.RegisterTranslation("mobile", s.Trans, func(ut ut.Translator) error {
		_ = v.RegisterTranslation("mobile", validators.InstanceTrans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 手机号码格式校验不通过！", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
	//
}
