package baidu_v1

import (
	"crypto/md5"
	"encoding/hex"
	"evlic.cn/bcp3th/lab/lab1/simple-dict/dict-api"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type BaiDuAPI struct {
}

func (b BaiDuAPI) Do(word string) dict_api.Result {
	realUrl := getUrl(word)
	log.Println("get Url >> ", realUrl)
	req, _ := http.NewRequest(http.MethodGet, realUrl, nil)
	resp, err := defClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return ReadResp(resp)
}

// ReadResp 响应处理
func ReadResp(resp *http.Response) dict_api.Result {
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(res))
	fmt.Println("read Body")
	return dict_api.Result{}
}

func md5Str(str string) string {
	h := md5.New()
	h.Write([]byte(str))

	MD5 := hex.EncodeToString(h.Sum(nil))
	log.Println("MD5 >> ", MD5)
	return MD5
}

func getUrl(word string) string {
	// salt := strconv.Itoa(rand.Intn(maxSalt-minSalt) + minSalt)
	salt := "1435660288"
	fmt.Println("salt >> ", salt)
	params := copyParams()
	params.Set("q", word)
	params.Set("salt", salt)
	params.Set("signer", md5Str(appId+word+salt+appKey))

	parseUrl, _ := url.Parse(baiduApiURL)
	parseUrl.RawQuery = params.Encode()
	return parseUrl.String()
}

const (
	appId       = "20220424001187650"
	appKey      = "tuyfqSfcDNudCbIiZDEJ"
	formLang    = "en"
	toLang      = "zh"
	baiduApiURL = "https://fanyi-api.baidu.com/api/trans/vip/translate"
	minSalt     = 0x800000
	maxSalt     = 0x1000000
)

var (
	// 会被改变，所以用法是 copy
	defParams = url.Values{
		"q":      []string{"hello"},
		"from":   []string{formLang},
		"to":     []string{toLang},
		"appid":  []string{appId},
		"salt":   []string{appId},
		"signer": []string{appId},
	}

	// defParseUrl, _ = url.Parse(baiduApiURL)
	defClient = http.DefaultClient
)

func copyParams() url.Values {
	res := url.Values{}
	for key, v := range defParams {
		for _, vv := range v {

			res.Add(key, vv)
		}
	}
	return res
}
