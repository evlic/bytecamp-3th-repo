package signer

import (
	"github.com/robertkrimen/otto"
	"log"
	"runtime"
)

var (
	p = "/"
)

func init() {

	if runtime.GOOS == "windows" {
		p = "\\"
	}
}

// Sign 返回百度的校验签名
func Sign(word string) string {
	js := getJS()
	vm := otto.New()
	_, err := vm.Run(js)
	if err != nil {
		log.Fatalln("run err!", err.Error())

	}
	res, err := vm.Call("e", nil, word)
	if err != nil {
		log.Fatalln("call err!", err.Error())
	}
	return res.String()
}

// 加载 js 脚本
func getJS() string {
	// pwd, _ := os.Getwd()
	// jsFile := "baiduv2.js"
	//
	// bytes, err := ioutil.ReadFile(pwd + p + jsFile)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// return string(bytes)
	return jsTOSign
}
