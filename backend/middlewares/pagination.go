package middlewares

import (
	"backend/models/form"
	"backend/pkg/constants"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/8/23 17:48
//

func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从请求中提取分页对象
		var pagination form.PaginateForm
		if err := c.ShouldBindQuery(&pagination); err == nil {
			// 解析参数
			if pagination.PageSize < 1 || pagination.PageSize > 100 {
				pagination.PageSize = 15
			}
			if pagination.PageNum < 1 {
				pagination.PageNum = 1
			}
			c.Set(constants.PAGINATION_KEY, pagination)
		}
		c.Next() // 放行继续执行后面的逻辑
	}
}
