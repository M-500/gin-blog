package jsons

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/8/22 11:00
//

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func Parse(str string, t interface{}) error {
	return json.Unmarshal([]byte(str), t)
}

func ToStr(t interface{}) (string, error) {
	if t == nil {
		return "", nil
	}
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ToJsonStr(t interface{}) string {
	str, err := ToStr(t)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	return str
}
