package main

// 注意，main方法必须在main包下，同一个包只能由一个名称。

import (
	"go_test/client"
	"go_test/model"
)

func main() {

	// 一个类型中的字段，可以重置，也可以使用默认值，在go中，所有的类型的初始值，都是字段类型的0值，比如string的初始值是""空字符串，int类型的初始值是0等等
	httpClientSettings := model.HttpClientSettings{
		Name:                     "测试厨房",
		NoDefaultUserAgentHeader: true,
		MaxConnDuration:          1000,
	}

	headers := []model.Header{
		model.Header{
			Field: "name",
			Value: "你好",
		},
	}

	httpRequest := model.HttpRequest{
		Name:               "planet",
		Url:                "http://www.baidu.com",
		Method:             "GET",
		HttpClientSettings: httpClientSettings,
		Headers:            headers,
	}

	client.RequestHttp(httpRequest)
}
