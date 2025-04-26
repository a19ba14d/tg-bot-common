package consts

import "github.com/gogf/gf/v2/frame/g"

type Language string

const (
	LangZhCN Language = "zh-CN"
	LangEn   Language = "en"
	// 添加其他支持的语言...

	// 默认语言
	DefaultLanguage = LangZhCN

	UserLanguageKeyPrefix = "user:lang:"
)

// DeptType 定义部门类型结构，方便API返回或前端使用
type LanguageType struct {
	Code Language `json:"code"` // 类型编码 (存储值)
	Name string   `json:"name"` // 类型名称 (显示值)
	Text string   `json:"text"` // 类型文本 (显示值)
}

var TypeList = g.Slice{
	LanguageType{Code: LangZhCN, Name: "中文", Text: "🇨🇳 简体中文"},
	LanguageType{Code: LangEn, Name: "English", Text: "🇺🇸 English"},
}

// TypeMap 提供 Code到Name 的快速查找映射
var TypeMap = func() g.MapStrStr {
	m := g.MapStrStr{}
	for _, v := range TypeList {
		item := v.(LanguageType) // 断言为 DeptType
		m[string(item.Code)] = item.Name
	}
	return m
}()

// 判断是否在TypeList中
func IsLanguageType(language Language) bool {
	for _, v := range TypeList {
		item := v.(LanguageType)
		if item.Code == language {
			return true
		}
	}
	return false
}
