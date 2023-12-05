package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go_test/model"
	"io/ioutil"
	"time"

	"github.com/valyala/fasthttp"
)

func RequestHttp(httpRequest model.HttpRequest) {

	// 使用fasthttp 协程池

	// 新建一个http请求
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	// 新建一个http响应接受服务端的返回
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 新建一个http的客户端
	client := newHttpClient(httpRequest.HttpClientSettings)

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
	//
	httpsTls := httpClientSettings.AdvancedOptions.Tls

	// 如果开启认证
	if httpsTls.IsVerify {
		// switch条件选择语句，如果认证类型为0：则表示双向认证，如果是1：则表示为单向认证
		switch httpsTls.VerifyType {
		case 0: // 开启双向验证
			tr.InsecureSkipVerify = false
			// 如果密钥文件为空则跳出switch语句
			if httpsTls.CaCert == "" {
				break
			}
			// 生成一个cert对象池
			caCertPool := x509.NewCertPool()
			if caCertPool == nil {
				fmt.Println("生成CertPool失败！")
				break
			}

			// 读取认证文件，读出后为字节数组
			key, err := ioutil.ReadFile(httpsTls.CaCert)
			// 如果读取错误，则跳出switch语句
			if err != nil {
				fmt.Println("打开密钥文件失败：", err.Error())
				break
			}
			// 将认证文件添加到cert池中
			ok := caCertPool.AppendCertsFromPEM(key)
			// 如果添加失败则跳出switch语句
			if !ok {
				fmt.Println("密钥文件错误，生成失败！！！")
				break
			}
			// 将认证信息添加到客户端认证结构体
			tr.ClientCAs = caCertPool
		case 1: // 开启单向验证，客户端验证服务端密钥
			tr.InsecureSkipVerify = false
		}
	}

	// 客户端认证配置项
	httpClient.TLSConfig = tr
	return
}
