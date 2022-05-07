package dict_api

import (
	"fmt"
)

// Result 英译中
type Result struct {
	// 音标
	En, Am string
	Means  []string
}

func (r Result) String() string {
	var wordMean string
	for _, mean := range r.Means {
		wordMean += mean
	}
	return fmt.Sprintf("英 %s\n美 %s\n\t%v", r.En, r.Am, r.Means)
}

type Translator interface {
	Do(word string) Result
}
