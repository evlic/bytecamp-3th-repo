package model

import (
	"encoding/json"
	dictApi "evlic.cn/bcp3th/lab/simple-dict/dict-api"
)

var (
	defRes = dictApi.Result{Means: []string{"查询出错"}}
)

func ParseToResult(jsonStr string) (dictApi.Result, error) {
	var obj Answer
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return defRes, err
	}
	s := obj.DictResult.Means.Symbols[0]

	return s.toResult(), nil
}
