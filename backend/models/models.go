package models

import (
	"gorm.io/gorm"
	"time"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/8/22 14:35
//

var Models = []interface{}{
	&Article{}, &ArticleTag{}, &Category{},
}

type BaseModel struct {
	ID        int            `gorm:"primary_key;auto_increment" json:"id" mapstructure:"id"`
	CreatedAt time.Time      `json:"created_at" mapstructure:"-"`
	UpdatedAt time.Time      `json:"updated_at" mapstructure:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 该字段不展示给前端
}
