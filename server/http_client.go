package server

import (
	"crypto/tls"
	"fmt"
	"go_test/model"
	"time"

	"github.com/valyala/fasthttp"
)

func RequestHttp(httpClientSettings model.HttpClientSettings) {

	// 使用fasthttp 协程池

	// 新建一个http请求
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	// 新建一个http响应接受服务端的返回
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 新建一个http的客户端
	client := newHttpClient(httpClientSettings)

	// 添加该请求的http方法：get、post、delete、update等等
	req.Header.SetMethod("GET")

	// 添加该请求的http的url
	req.SetRequestURI("http://www.baidu.com")

	// 开始请求
	err := client.Do(req, resp)
	if err != nil {
		fmt.Sprintln("发送http请求错误：", err.Error())
	}

	// fmt.Println("resp:    ", resp.String())
	fmt.Println("resp:    ", req.Header.Header())

}

func newHttpClient(httpClientSettings model.HttpClientSettings) (httpClient *fasthttp.Client) {
	// tls验证，关闭验证
	tr := &tls.Config{
		InsecureSkipVerify: true,
	}
	// 新建指针类型的客户端
	httpClient = &fasthttp.Client{}
	httpClient.TLSConfig = tr
	if httpClientSettings.Name != "" {
		httpClient.Name = httpClientSettings.Name
	}
	if httpClientSettings.NoDefaultUserAgentHeader == true {
		httpClient.NoDefaultUserAgentHeader = true
	}
	// 如果最大连接数不为0，将设置此数
	if httpClientSettings.MaxConnsPerHost != 0 {
		httpClient.MaxConnsPerHost = httpClientSettings.MaxConnsPerHost
	}
	// url不按照标准输出，按照原样输出
	if httpClientSettings.DisablePathNormalizing == true {
		httpClient.DisablePathNormalizing = true
	}
	// 请求头不按标准格式传输
	if httpClientSettings.DisableHeaderNamesNormalizing == true {
		httpClient.DisableHeaderNamesNormalizing = true
	}
	// 如果此时间不为0，那么将设置此时间。keep-alive维持此时长后将关闭。时间单位为毫秒
	if httpClientSettings.MaxConnDuration != 0 {
		httpClient.MaxConnDuration = time.Duration(httpClientSettings.MaxConnDuration) * time.Millisecond
	}
	if httpClientSettings.ReadTimeout != 0 {
		httpClient.ReadTimeout = time.Duration(httpClientSettings.ReadTimeout) * time.Millisecond
	}
	if httpClientSettings.WriteTimeout != 0 {
		httpClient.WriteTimeout = time.Duration(httpClientSettings.WriteTimeout) * time.Millisecond
	}
	// 该连接如果空闲的话，在此时间后断开。
	if httpClientSettings.MaxIdleConnDuration != 0 {
		httpClient.MaxIdleConnDuration = time.Duration(httpClientSettings.MaxIdleConnDuration) * time.Millisecond
	}
	return
}
