package baidu_v2

import (
	"evlic.cn/bcp3th/lab/lab1/simple-dict/dict-api"
	"evlic.cn/bcp3th/lab/lab1/simple-dict/dict-api/baidu-v2/model"
	"evlic.cn/bcp3th/lab/lab1/simple-dict/dict-api/baidu-v2/signer"
	"io"
	"log"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"
)

const (
	DefUrl = "https://fanyi.baidu.com/v2transapi"
)

var (
	// 会被改变，所以用法是 copy
	defParams = url.Values{
		"from": []string{"en"},
		"to":   []string{"zh"},
	}

	defValue = url.Values{
		"from": []string{"en"},
		"to":   []string{"zh"},
		// "query":             []string{"xx"},
		// "sign":              []string{"932273.711296"},
		"transtype":         []string{"translang"},
		"simple_means_flag": []string{"3"},
		"token":             []string{"748d93189f5da71831b954e218705f8f"},
		"domain":            []string{"common"},
	}

	defHeader = http.Header{
		textproto.CanonicalMIMEHeaderKey(`accept`):       []string{`*/*`},
		textproto.CanonicalMIMEHeaderKey(`Content-Type`): []string{`application/x-www-form-urlencoded; charset=UTF-8`},
		textproto.CanonicalMIMEHeaderKey(`Connection`):   []string{`keep-alive`},
		// textproto.CanonicalMIMEHeaderKey(`accept-encoding`):    []string{`gzip`, `deflate`, `br`},
		textproto.CanonicalMIMEHeaderKey(`accept-language`):    []string{`zh`, `zh-CN;q=0.9`, `en-US;q=0.8`, `en;q=0.7`},
		textproto.CanonicalMIMEHeaderKey(`cookie`):             []string{`BAIDUID=E227F1E07273FFE2560F809666E0EFEA:FG=1; BAIDU_WISE_UID=wapp_1650777977487_527; BIDUPSID=E227F1E07273FFE2560F809666E0EFEA; PSTM=1650808308; BAIDUID_BFESS=C4FD1C48F20E8D695CEB64AB84D638AF:FG=1; MCITY=-132:; REALTIME_TRANS_SWITCH=1; FANYI_WORD_SWITCH=1; HISTORY_SWITCH=1; SOUND_SPD_SWITCH=1; SOUND_PREFER_SWITCH=1; BDUSS=dHTmNRMUVJM21LeUtMbW5OQ0xPdjhSMnBVeDNOZW0tSUFuODQzekpreU14NXhpRUFBQUFBJCQAAAAAAAAAAAEAAABuuugmyfrLwNeqwtYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIw6dWKMOnVidj; BDUSS_BFESS=dHTmNRMUVJM21LeUtMbW5OQ0xPdjhSMnBVeDNOZW0tSUFuODQzekpreU14NXhpRUFBQUFBJCQAAAAAAAAAAAEAAABuuugmyfrLwNeqwtYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIw6dWKMOnVidj; RT="z=1&dm=baidu.com&si=vf4jnmh086&ss=l2ultkkc&sl=2&tt=53d&bcn=https://fclog.baidu.com/log/weirwood?type=perf&ld=3hc&ul=48v&hd=498"; ZD_ENTRY=bing; Hm_lvt_64ecd82404c51e03dc91cb9e8c025574=1651824455,1651853006; Hm_lpvt_64ecd82404c51e03dc91cb9e8c025574=1651853006; ab_sr=1.0.1_NDQ0N2I5MTQ2N2M3MGFmNjMwODk5M2IyZTQ2MDEyYzg5NmY1YjYzN2I2MWZhOTUzMThmNTNmZmQxNTcwNDhjMDAyYzlhZjMxYTIwNWVjNDU1MjBiMGJkMjk3NGUzY2Q5YzZlM2U3YWQ0ZDFmODM2ODA3YzRlNWZiNWQwYmQxMmE3MGY5MTEwNTQxMTI3NmYyYTUxZDhmMDU1MzRhZjU3ZDIyNmQ0ZWU1ZDlkZGQwMWRmOGY4NTU3YTY1MzBiYTRl`},
		textproto.CanonicalMIMEHeaderKey(`referer`):            []string{`https://fanyi.baidu.com/`},
		textproto.CanonicalMIMEHeaderKey(`sec-ch-ua`):          []string{`" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100`},
		textproto.CanonicalMIMEHeaderKey(`sec-ch-ua-mobile`):   []string{`?0`},
		textproto.CanonicalMIMEHeaderKey(`sec-ch-ua-platform`): []string{`"Windows"`},
		textproto.CanonicalMIMEHeaderKey(`sec-fetch-dest`):     []string{`empty`},
		textproto.CanonicalMIMEHeaderKey(`sec-fetch-mode`):     []string{`cors`},
		textproto.CanonicalMIMEHeaderKey(`sec-fetch-site`):     []string{`same-origin`},
		textproto.CanonicalMIMEHeaderKey(`traceparent`):        []string{`00-b62ed7fb49dc5482684c90c1a4a45407-4bcc539a429b6960-01`},
		textproto.CanonicalMIMEHeaderKey(`user-agent`):         []string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36`},
	}

	// defParseUrl, _ = url.Parse(DefUrl)
	defClient = http.DefaultClient
)

type BaiDuAPI struct {
}

func (b BaiDuAPI) Do(s string) dict_api.Result {
	urlTarget := getUrl()
	val := copyValues()
	val.Set("query", s)
	val.Set("sign", signer.Sign(s))

	req := NewReq(urlTarget, strings.NewReader(val.Encode()))
	resp, _ := defClient.Do(req)
	readResp := ReadResp(resp)
	// fmt.Println("json >> ", readResp)
	result, err := model.ParseToResult(readResp)
	if err != nil {
		log.Fatalln(err)
	}
	return result
}

func getUrl() string {
	params := copyParams()
	parseUrl, _ := url.Parse(DefUrl)
	parseUrl.RawQuery = params.Encode()
	return parseUrl.String()
}

// NewReq 封装默认 req 设置好了 defHeader
func NewReq(reqUrl string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, reqUrl, body)
	req.Header = defHeader
	return req
}
