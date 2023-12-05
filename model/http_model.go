package model

import (
	"crypto/tls"
)

// HttpRequest http请求的结构
type HttpRequest struct {
	Name               string             // 接口名称
	Url                string             // 接口uri
	Method             string             // 接口方法，Get Post Update...
	Headers            []Header           // 接口请求头
	Querys             []Query            // get请求时的url
	Cookies            []Cookie           // cookie
	Body               string             // 请求体
	HttpClientSettings HttpClientSettings // http客户端配置
}

// Header header
type Header struct {
	Field     string // 字段名称
	Value     string // 字段值
	FieldType string // 字段类型
}

// Query query
type Query struct {
	Field     string
	Value     string
	FieldType string
}

// Cookie cookie
type Cookie struct {
	Field     string
	Value     string
	FieldType string
}

type HttpClientSettings struct {
	//  客户端的名称，在header中的user-agent使用，通常我们默认就好
	Name string
	// 默认为flase，表示User-Agent使用fasthttp的默认值
	NoDefaultUserAgentHeader bool
	// https连接的TLS配置。这里使用的是tls.Config指针类型。可以在我们使用的时候配置
	TLSConfig *tls.Config
	// 每台主机可以建立的最大连接数。如果没有设置，则使用DefaultMaxConnsPerHost。
	MaxConnsPerHost int
	// 空闲的保持连接在此持续时间之后关闭。默认情况下，在DefaultMaxIdleConnDuration之后关闭空闲连接。
	// 该连接如果空闲的话，在此时间后断开。
	MaxIdleConnDuration int64
	// Keep-alive连接在此持续时间后关闭。默认情况下，连接时间是不限制的。
	MaxConnDuration int64
	// 默认情况下，响应读取超时时间是不限制的。
	ReadTimeout int64
	// 默认情况下，请求写超时时间不受限制。
	WriteTimeout int64
	// 请求头是否按标准格式传输
	DisableHeaderNamesNormalizing bool
	// url路径是按照原样输出，还是按照规范化输出。默认按照规范化输出
	DisablePathNormalizing bool
}
