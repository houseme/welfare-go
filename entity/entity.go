package entity

// Response 响应
type Response struct {
	Code    int         `json:"code" dc:"Error code"`                                                          // 响应代码
	Time    int64       `json:"time" dc:"Time of response"`                                                    // 时间戳
	Data    interface{} `json:"data,omitempty"  dc:"Result data for certain request according API definition"` // 业务域 加密方式:数据进行请求报文 encryptType 字段指定算法加密后并转码为 16 进制;
	Message string      `json:"message" dc:"Error message"`                                                    // 响应描述
	Sign    string      `json:"sign" dc:"sign"`                                                                // 数据签名域 由请求报文的 signType 指定签名方式生成签 名，参考后续安全要求的签名章节
	TraceId string      `json:"traceId"  dc:"Trace ID"`                                                        // 流水号
}

// Request 请求
type Request struct {
	Version     string `v:"required|in:2.0.0" json:"version" dc:"版本信息"`       // 2.0.0 默认版本
	Ts          string `v:"required|size:14" json:"ts" dc:"当前时间"`             // 格式:YYYYMMDDH24MISS
	Channel     string `v:"required|min-length:6" json:"channel" dc:"渠道唯一标识"` // 分配商户的唯一编号;
	NonceStr    string `v:"required|length:24,32" json:"nonceStr" dc:"随机字符串"` // 随机字符串，可参考随机字符串生产算法
	Format      string `v:"required|in:json" json:"format" dc:"数据格式仅支持 json"` // 返回的数据格式仅支持 json
	EncryptType string `v:"required|in:RSA2" json:"encryptType" dc:"数据加密类型"`  // 参考安全要求 RSA2
	Data        string `v:"required" json:"data" dc:"请求参数的集合"`                // 请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体该字段必须加密加密传输;加密方式:为 json 数据进行 encryptType 字 段指定算法加密后并转码为 16 进制;
	SignType    string `v:"required|in:RSA2" json:"signType" dc:"签名类型"`       // 参考安全要求 rsa2
	Sign        string `v:"required" json:"sign" dc:"数据签名域"`                  // 数据签名域 由 signType指定签名方式生成签名，参考后续 安全要求的签名章节
	ServiceType string `v:"required" json:"serviceType" dc:"服务类型"`            // 服务类型
}

// CreateUserInput 创建用户输入
type CreateUserInput struct {
	Username string `json:"username" dc:"用户名"`
	Mobile   string `json:"mobile" dc:"手机号"`
	Number   string `json:"number" dc:"用户唯一编号"`
}

// CreateUserOutput 创建用户输出
type CreateUserOutput struct {
	OpenID string `json:"openId" dc:"用户唯一标识"`
}

// AuthInput is the input body for the /auth endpoint
type AuthInput struct {
	OpenID string `json:"openId"`
}

// AuthOutput is the output body for the /auth endpoint
type AuthOutput struct {
	Schema string `json:"schema"`
}
