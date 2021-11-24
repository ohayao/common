package file

import (
	"encoding/json"
	"io/ioutil"
)

// GetJSONConfig 获取配置文件到结构体 jsonPath 文件路径 result 指针类型
func GetJSONConfig(jsonPath string, result interface{}) error {
	data, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return err
	}
	return nil
}

// GetFileContent 获取文件内容
func GetFileContent(filePath string) string {
	if data, err := ioutil.ReadFile(filePath); err != nil {
		return ""
	} else {
		return string(data)
	}
}
