package form

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/8/23 17:49
//
type PaginateForm struct {
	PageSize int64 `json:"pageSize" form:"pageSize" binding:"-"`
	PageNum  int64 `json:"pageNum" form:"pageNum" binding:"-"`
}
