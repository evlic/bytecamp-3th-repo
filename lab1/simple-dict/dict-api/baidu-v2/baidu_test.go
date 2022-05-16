package baidu_v2

import (
	"evlic.cn/bcp3th/lab/simple-dict/dict-api/baidu-v2/signer"
	"fmt"
	"testing"
)

var (
	bd BaiDuAPI
)

func TestBaiDuAPIV2(t *testing.T) {
	fmt.Println(bd.Do("hello"))
}

func TestSign(t *testing.T) {
	s := signer.Sign("gopher")
	fmt.Println(s)
}
