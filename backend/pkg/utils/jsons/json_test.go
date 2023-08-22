package jsons

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/8/22 11:00
//

import (
	"fmt"
	"testing"
)

func TestToJsonStr(t *testing.T) {
	ret := ToJsonStr(nil)
	fmt.Println(ret)
	fmt.Println(ret == "")
	if ret != "" {
		t.Fail()
	}
}
