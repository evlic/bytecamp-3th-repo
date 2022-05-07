package baidu_v2

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func ReadResp(resp *http.Response) string {
	defer resp.Body.Close()
	// 默认使用 gzip
	all, _ := ioutil.ReadAll(resp.Body)
	return string(all)
}

func copyUrlValues(u url.Values) url.Values {
	res := url.Values{}
	for key, v := range u {
		for _, vv := range v {
			res.Add(key, vv)
		}
	}
	return res
}

func copyParams() url.Values {
	return copyUrlValues(defParams)
}

func copyValues() url.Values {
	return copyUrlValues(defValue)
}
