package gtranslate

import (
	"net/http"
	"os"
	"testing"
	"time"
)

func TestTranslateWithFromTo(t *testing.T) {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:41091")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:41091")
	httpClient := http.Client{
		Transport: &http.Transport{
			// 设置代理，从环境变量中获取
			Proxy: http.ProxyFromEnvironment,
		},
	}
	for i := 0; i < 1; i++ {
		for _, ta := range testingTable {
			resp, err := TranslateWithParams(ta.inText, TranslationParams{
				From:       ta.langFrom,
				To:         ta.langTo,
				Tries:      5,
				Delay:      time.Second,
				GoogleHost: "google.cn",
				HttpClient: httpClient,
			})
			if err != nil {
				t.Error(err, err.Error())
				t.Fail()
			}
			println(ta.inText, "-", resp)
			if resp != ta.outText {
				t.Error("translated text is not the expected", ta.outText, " != ", resp)
			}
		}
	}
}
