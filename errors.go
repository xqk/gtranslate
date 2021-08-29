package gtranslate

import "errors"

var errBadNetwork = errors.New("网络故障，请检查您的网络连接")
var errBadRequest = errors.New("错误请求，谷歌翻译API上的请求不生效")
