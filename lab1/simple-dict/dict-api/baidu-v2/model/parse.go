package model

import (
	"encoding/json"
	dict_api "evlic.cn/bcp3th/lab/lab1/simple-dict/dict-api"
)

var (
	defRes = dict_api.Result{Means: []string{"查询出错"}}
)

func ParseToResult(jsonStr string) (dict_api.Result, error) {
	var obj Answer
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return defRes, err
	}
	s := obj.DictResult.Means.Symbols[0]

	return s.toResult(), nil
}
