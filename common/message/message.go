package message

// 定义一些消息类型
const (
	LoginMessageType     = "LoginMessage"
	LoginResultType      = "LoginResultType"
	RegisterMessageType  = "RegisterMessageType"
	RegisterResultType   = "RegisterResultType"
	NotifyUserStatusType = "NotifyUserStatus"
	SmsMessageType       = "SmsMessageType"
	SmsResultType        = "SmsResultType"
	SmsRecordType        = "SmsRecordType"
	SmsRecodeResultType  = "SmsRecodeResultType"
)

const (
	UserOnline = iota
	UserOffline
)

type (
	Message struct {
		Type string `json:"type"` // 消息类型
		Data string `json:"data"` // 消息内容
	}

	LoginMessage struct {
		UserId   int    `json:"user_id"`   // 用户id
		UserPwd  string `json:"user_pwd"`  // 用户密码
		UserName string `json:"user_name"` // 用户名
	}

	LoginResult struct {
		Code  int    `json:"code"`            // 返回状态码 200成功 444未注册
		Error string `json:"error"`           // 返回错误信息
		Users []User `json:"users,omitempty"` // 存放所有在线用户信息
	}

	RegisterMessage struct {
		UserId   int    `json:"user_id,omitempty"`
		UserPwd  string `json:"user_pwd,omitempty"`
		UserName string `json:"user_name,omitempty"`
	}

	RegisterResult struct {
		Code  int    `json:"code,omitempty"`  // 返回状态码 200成功 444已被注册
		Error string `json:"error,omitempty"` // 返回错误信息
	}

	User struct {
		UserId   int    `json:"user_id,omitempty"`   // 用户id
		UserName string `json:"user_name,omitempty"` // 用户名
		Status   int    `json:"status,omitempty"`    // 用户状态
	}

	SmsMessage struct {
		Content string        `json:"content,omitempty"` // 发送内容
		User    `json:"user"` // 用户信息
	}

	SmsRecord struct {
		Records []SmsMessage
	}
)
