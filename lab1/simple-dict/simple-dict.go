package main

import (
	api "evlic.cn/bcp3th/lab/simple-dict/dict-api"
	baiduV2 "evlic.cn/bcp3th/lab/simple-dict/dict-api/baidu-v2"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Println("请输入要翻译的单词")
		return
	}
	word := os.Args[1]
	var (
		apis = []api.Translator{
			baiduV2.BaiDuAPI{},
			api.ColorCloudAPI{},
		}
	)
	res := make(chan api.Result)
	defer close(res)

	for _, tApi := range apis {
		go func(t api.Translator) {
			res <- presses(t, word)
		}(tApi)
	}
	log.Printf("\n%v", <-res)
}

func presses(translator api.Translator, word string) api.Result {
	return translator.Do(word)
}
