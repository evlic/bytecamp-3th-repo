package model

import (
	dict_api "evlic.cn/bcp3th/lab/simple-dict/dict-api"
)

func (s Symbol) toResult() dict_api.Result {
	means := make([]string, 0, len(s.Parts))
	for _, part := range s.Parts {
		var str string
		str += part.Part

		for idx, mean := range part.Means {
			str += mean
			if idx != len(part.Means)-1 {
				str += ", "
			} else {
				str += "; "
			}
		}

		means = append(means, str)
	}

	return dict_api.Result{
		En:    s.PhEn,
		Am:    s.PhAm,
		Means: means,
	}
}
