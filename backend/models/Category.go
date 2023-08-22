package models

import "reflect"

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/8/22 14:55
//

// hasMany: 一个分类下可以有多篇文章
type Category struct {
	BaseModel
	Name     string    `gorm:"type:varchar(20);not null;comment:分类名称" json:"name"`
	Articles []Article `gorm:"foreignKey:CategoryId"` // 重写外键
}

func (c *Category) IsEmpty() bool {
	return reflect.DeepEqual(c, &Category{})
}
