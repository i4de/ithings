// Code generated by goctl. DO NOT EDIT.
package types

type GetCaptchaReq struct {
	Data string `json:"data,optional"`              //短信验证时填写手机号
	Type string `json:"type,options=sms|image"`     //验证方式:短信验证,图片验证码
	Use  string `json:"use,options=login|register"` //用途
}

type GetCaptchaResp struct {
	CodeID string `json:"codeID"`       //验证码编号
	Url    string `json:"url,optional"` //图片验证码url
	Expire int64  `json:"expire"`       //过期时间
}

type PageInfo struct {
	Page       int64  `json:"page,optional" form:"page,optional"`             // 页码
	PageSize   int64  `json:"pageSize,optional" form:"pageSize,optional"`     // 每页大小
	SearchKey  string `json:"searchKey,optional" form:"searchKey,optional"`   // 搜索的key
	SearchType string `json:"searchType,optional" form:"searchType,optional"` // 搜索的类型
}

type JwtToken struct {
	AccessToken  string `json:"accessToken,omitempty"`
	AccessExpire int64  `json:"accessExpire,string,omitempty"`
	RefreshAfter int64  `json:"refreshAfter,string,omitempty"`
}

type UserInfo struct {
	Uid        int64  `json:"uid,string"`                 // 用户id
	UserName   string `json:"userName,optional"`          //用户名(唯一)
	NickName   string `json:"nickName,optional"`          // 用户的昵称
	InviterUid int64  `json:"inviterUid,string,optional"` // 邀请人用户id
	InviterId  string `json:"inviterId,optional"`         // 邀请码
	Sex        int64  `json:"sex,optional"`               // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City       string `json:"city,optional"`              // 用户所在城市
	Country    string `json:"country,optional"`           // 用户所在国家
	Province   string `json:"province,optional"`          // 用户所在省份
	Language   string `json:"language,optional"`          // 用户的语言，简体中文为zh_CN
	HeadImgUrl string `json:"headImgUrl,optional"`        // 用户头像
	CreateTime int64  `json:"createTime,string,optional"`
}

type LoginReq struct {
	UserID    string `json:"userID,optional"`                               //登录账号(支持用户名,手机号登录) 账号密码登录时需要填写
	PwdType   int32  `json:"pwdtype,optional"`                              //账号密码登录时需要填写.0,无密码 1，明文 2，md5加密
	Password  string `json:"password,optional"`                             //密码，建议md5转换 密码登录时需要填写
	LoginType string `json:"loginType,options=sms|img|wxopen|wxin|wxminip"` //验证类型 sms 短信验证码 img 图形验证码加账号密码登录 wxopen 微信开放平台登录 wxin 微信内登录 wxminip 微信小程序
	Code      string `json:"code,optional"`                                 //验证码    微信登录填code
	CodeID    string `json:"codeID,optional"`                               //验证码编号 微信登录填state
}

type LoginResp struct {
	Info  UserInfo `json:"info"`  //用户信息
	Token JwtToken `json:"token"` //用户token
}

type UserCore struct {
	Uid         int64  `json:"uid,string"`
	UserName    string `json:"userName"`
	Password    string `json:"password,omitempty"` // 登录密码
	Email       string `json:"email"`              // 邮箱
	Phone       string `json:"phone"`              // 手机号
	Wechat      string `json:"wechat"`             // 微信UnionID
	LastIP      string `json:"lastIP"`             // 最后登录ip
	RegIP       string `json:"regIP"`              // 注册ip
	CreatedTime int64  `json:"createdTime,string"` //创建时间
	Status      int64  `json:"status"`             // 用户状态:0为未注册状态
}

type RegisterCoreReq struct {
	ReqType string `json:"reqType,options=phone|wxopen|wxin|wxminip"` //注册方式:	phone手机号注册 wxopen 微信开放平台登录 wxin 微信内登录 wxminip 微信小程序
	Note    string `json:"note,optional"`                             //手机号注册时填写手机号
	Code    string `json:"code"`                                      //验证码    微信登录填code
	CodeID  string `json:"codeID,optional"`                           //验证码编号 微信登录填state
}

type RegisterCoreResp struct {
	Uid int64 `json:"uid,string"` //用户id
	JwtToken
}

type Register2Req struct {
	Token    string `json:"token"`             //注册第一步的token
	Password string `json:"password,optional"` //明文密码
	UserInfo
}

type ModifyStu struct {
	Type  string `json:"type"`  //修改类型
	Value string `json:"value"` //修改的值
}

type ModifyUserInfoReq struct {
	Info map[string]string `json:"info"` //修改参数key value数组
}

type GetUserCoreListReq struct {
	PageInfo
}

type GetUserCoreListResp struct {
	Total int64       `json:"total"`
	Info  []*UserCore `json:"info"`
}

type GetUserInfosReq struct {
	Uid []string `json:"uid,string"` //十进制数
}

type GetUserInfosResp struct {
	Info []*UserInfo `json:"info"`
}

type LoginAuthReq struct {
	Username    string `json:"username"`                       //用户名
	Password    string `json:"password,optional"`              //密码
	ClientID    string `json:"clientID"`                       //clientID
	Ip          string `json:"ip"`                             //访问的ip地址
	Certificate string `json:"certificate,optional,omitempty"` //客户端证书 base64后传过来
}

type AccessAuthReq struct {
	Username string `json:"username,omitempty"` //用户名
	Topic    string `json:"topic,omitempty"`    //主题
	ClientID string `json:"clientID,omitempty"` //clientID
	Access   string `json:"access,omitempty"`   //操作
	Ip       string `json:"ip,omitempty"`       //访问的ip地址
}

type RootCheckReq struct {
	Username    string `json:"username,omitempty"`             //用户名
	Password    string `json:"password,optional,omitempty"`    //密码
	ClientID    string `json:"clientID,omitempty"`             //clientID
	Ip          string `json:"ip,omitempty"`                   //访问的ip地址
	Certificate []byte `json:"certificate,optional,omitempty"` //客户端证书
}

type ProductTemplate struct {
	CreatedTime int64  `json:"createdTime,optional,string,omitempty"` //创建时间 只读
	ProductID   string `json:"productID,optional,omitempty"`          //产品id 只读
	Template    string `json:"template,optional,omitempty"`           //数据模板
}

type ProductInfo struct {
	CreatedTime  int64   `json:"createdTime,optional,string,omitempty"` //创建时间 只读
	ProductID    string  `json:"productID,optional,omitempty"`          //产品id 只读
	ProductName  string  `json:"productName,optional,omitempty"`        //产品名称
	AuthMode     int64   `json:"authMode,optional"`                     //认证方式:1:账密认证,2:秘钥认证
	DeviceType   int64   `json:"deviceType,optional"`                   //设备类型:1:设备,2:网关,3:子设备
	CategoryID   int64   `json:"categoryID,optional"`                   //产品品类
	NetType      int64   `json:"netType,optional"`                      //通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN
	DataProto    int64   `json:"dataProto,optional"`                    //数据协议:1:自定义,2:数据模板
	AutoRegister int64   `json:"autoRegister,optional"`                 //动态注册:1:关闭,2:打开,3:打开并自动创建设备
	Secret       string  `json:"secret,optional,omitempty"`             //动态注册产品秘钥 只读
	Description  *string `json:"description,optional,omitempty"`        //描述
	DevStatus    *string `json:"devStatus,optional,omitempty"`          // 产品状态
}

type ManageProductReq struct {
	Opt  int64        `json:"opt,omitempty"`           //操作 0:新增 1:修改 2:删除
	Info *ProductInfo `json:"info,optional,omitempty"` //产品信息
}

type GetProductInfoReq struct {
	ProductID string    `json:"productID,optional,optional"` //产品id  如果为空返回所有产品的信息
	Page      *PageInfo `json:"page,optional,optional"`      //分页信息,只获取一个则不填
}

type GetProductInfoResp struct {
	List  []*ProductInfo `json:"list,omitempty"`           //产品信息
	Total int64          `json:"total,optional,omitempty"` //拥有的总数(只有分页的时候会返回)
	Num   int64          `json:"num,optional,omitempty"`   //返回的数量
}

type ManageProductTemplateReq struct {
	Info *ProductTemplate `json:"info,optional,omitempty"` //产品信息
}

type GetProductTemplateReq struct {
	ProductID string `json:"productID"` //产品id
}

type GetProductTemplateResp struct {
	Info ProductTemplate `json:"info,omitempty"` //产品信息
}

type DeviceInfo struct {
	ProductID   string  `json:"productID"`                   //产品id 只读
	DeviceName  string  `json:"deviceName"`                  //设备名称 读写
	CreatedTime int64   `json:"createdTime,optional,string"` //创建时间 只读
	Secret      string  `json:"secret,optional"`             //设备秘钥 只读
	FirstLogin  int64   `json:"firstLogin,optional,string"`  //激活时间 只读
	LastLogin   int64   `json:"lastLogin,optional,string"`   //最后上线时间 只读
	Version     *string `json:"version,optional"`            // 固件版本  读写
	LogLevel    int64   `json:"logLevel,optional"`           // 日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写
	Cert        string  `json:"cert,optional,omitempty"`     // 设备证书  只读
}

type ManageDeviceReq struct {
	Opt  int64       `json:"opt,omitempty"`  //操作 0:新增 1:修改 2:删除
	Info *DeviceInfo `json:"info,omitempty"` //设备信息
}

type GetDeviceInfoReq struct {
	DeviceName string    `json:"deviceName,optional"` //设备名 为空时获取产品id下的所有设备信息
	ProductID  string    `json:"productID,omitempty"` //产品id
	Page       *PageInfo `json:"page,optional"`       //分页信息 只获取一个则不填
}

type GetDeviceInfoResp struct {
	List  []*DeviceInfo `json:"list,omitempty"`  //设备信息
	Total int64         `json:"total,omitempty"` //总数(只有分页的时候会返回)
	Num   int64         `json:"num,omitempty"`   //返回的数量
}

type SendDmPropertyReq struct {
	ProductID     string `json:"productID,omitempty"`                     //产品id 获取产品id下的所有设备信息
	DeviceName    string `json:"deviceName,omitempty"`                    //设备名
	Data          string `json:"data,omitempty"`                          //属性数据, JSON格式字符串, 注意字段需要在物模型属性里定义
	DataTimestamp int64  `json:"dataTimestamp,string,omitempty,optional"` //上报数据UNIX时间戳, 仅对Method:reported有效
	Method        string `json:"method,omitempty,optional"`               //请求类型 , 不填该参数或者 desired 表示下发属性给设备, reported 表示模拟设备上报属性
}

type SendDmPropertyResp struct {
	ClientToken string `json:"clientToken"`    //调用id
	Status      string `json:"status"`         //返回状态
	Code        int64  `json:"code,string"`    //设备返回状态码
	Data        string `json:"data,omitempty"` //返回信息
}

type SendDmActionReq struct {
	ProductID   string `json:"productID,omitempty"`   //产品id 获取产品id下的所有设备信息
	DeviceName  string `json:"deviceName,omitempty"`  //设备名
	ActionId    string `json:"actionId,omitempty"`    //产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	InputParams string `json:"inputParams,omitempty"` //输入参数
}

type SendDmActionResp struct {
	ClientToken  string `json:"clientToken,omitempty"`  //调用id
	OutputParams string `json:"outputParams,omitempty"` //输出参数 注意：此字段可能返回 null，表示取不到有效值。
	Status       string `json:"status,omitempty"`       //返回状态
	Code         int64  `json:"code,string,omitempty"`  //设备返回状态码
}

type GetDeviceDataReq struct {
	Method     string `form:"method,omitempty"`                    //获取的类型   property 属性 event事件  log 所有日志
	DeviceName string `form:"deviceName,omitempty"`                //设备名
	ProductID  string `form:"productID,omitempty"`                 //产品id 获取产品id下的所有设备信息
	DataID     string `form:"dataID,optional,omitempty"`           //获取的具体标识符的数据 如果不指定则获取所有属性数据,一个属性一条,如果没有获取到的不会返回值
	TimeStart  int64  `form:"timeStart,string,optional,omitempty"` //获取时间的开始
	TimeEnd    int64  `form:"timeEnd,string,optional,omitempty"`   //时间的结束
	Limit      int64  `form:"limit,optional,omitempty"`            //获取的长度
}

type DeviceData struct {
	Timestamp int64  `json:"timestamp,string"`    //发生时间戳
	Type      string `json:"type,omitempty"`      //类型 事件类型: 信息:info  告警alert  故障:fault
	DataID    string `json:"dataID"`              //获取的具体属性值
	GetValue  string `json:"getValue,omitempty"`  //获取到的值
	SendValue string `json:"sendValue,omitempty"` //发送过去的参数(action限定)
}

type GetDeviceDataResp struct {
	Total int64         `json:"total,string"` //总数
	List  []*DeviceData `json:"list"`         //总数
}

type GetDeviceDescribeLogReq struct {
	ProductID  string `form:"productID,omitempty"`
	DeviceName string `form:"deviceName,omitempty"`
	TimeStart  int64  `form:"timeStart,optional,string,omitempty"`
	TimeEnd    int64  `form:"timeEnd,optional,string,omitempty"`
	Limit      int64  `form:"limit,optional,omitempty"`
}

type GetDeviceDescribeLogResp struct {
	List []*DeviceDescribeLog `json:"list,omitempty"`
}

type DeviceDescribeLog struct {
	Timestamp  int64  `json:"timestamp,string"`
	Action     string `json:"action"`
	RequestID  string `json:"requestID"`
	TranceID   string `json:"tranceID"`
	Topic      string `json:"topic"`
	Content    string `json:"content"`
	ResultType int64  `json:"resultType,string"`
}

type GetDevicePropertyStatusReq struct {
	ProductID  string `form:"productID,omitempty"`
	DeviceName string `form:"deviceName,omitempty"`
}

type DevicePropertyStatus struct {
	Type  string  `json:"type"`
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Time  int64   `json:"time,string"`
	Unit  string  `json:"unit"`
	Value *string `json:"value"`
}

type GetDevicePropertyStatusResp struct {
	List []*DevicePropertyStatus `json:"list"`
}

type GroupInfo struct {
	GroupID     int64  `json:"groupID,string,omitempty"`              //组id
	Name        string `json:"name,omitempty"`                        //组名
	Uid         int64  `json:"uid,string,omitempty"`                  //管理员用户id
	CreatedTime int64  `json:"createdTime,string,optional,omitempty"` //创建时间 只读
}

type GroupMember struct {
	GroupID     int64  `json:"groupID,omitempty"`                     //组id
	MemberID    string `json:"memberID,omitempty"`                    //成员id
	MemberType  int64  `json:"memberType,omitempty"`                  //成员类型:1:设备 2:用户
	CreatedTime int64  `json:"createdTime,string,optional,omitempty"` //创建时间 只读
}

type ManageGroupInfoReq struct {
	Opt  int64      `json:"opt,omitempty"`           //操作 0:新增 1:修改 2:删除
	Info *GroupInfo `json:"info,optional,omitempty"` //组信息
}

type ManageGroupMemberReq struct {
	Opt  int64        `json:"opt,omitempty"`           //操作 0:新增 1:修改 2:删除
	Info *GroupMember `json:"info,optional,omitempty"` //组成员
}

type GetGroupInfoReq struct {
	Page    *PageInfo `json:"page,optional,omitempty"`           //分页信息,只获取一个则不填
	GroupID int64     `json:"groupID,string,optional,omitempty"` //产品id  如果为空返回所有产品的信息
}

type GetGroupInfoResp struct {
	Info  []*GroupInfo `json:"info,omitempty"`           //产品信息
	Num   int64        `json:"num,omitempty"`            //返回的数量
	Total int64        `json:"total,optional,omitempty"` //拥有的总数(只有分页的时候会返回)
}

type GetGroupMemberReq struct {
	Page       *PageInfo `json:"page,optional,omitempty"`           //分页信息,两个参数都填不用填分页信息,用来确认组内是否有该成员
	GroupID    int64     `json:"groupID,optional,string,omitempty"` //组id
	MemberID   string    `json:"memberID,optional,omitempty"`       //成员id
	MemberType int64     `json:"memberType,optional,omitempty"`     //成员类型:1:设备 2:用户
}

type GetGroupMemberResp struct {
	Info  []*GroupMember `json:"info,omitempty"`           //产品信息
	Num   int64          `json:"num,optional,omitempty"`   //返回的数量
	Total int64          `json:"total,optional,omitempty"` //拥有的总数(只有分页的时候会返回)
}

type SendDcActionReq struct {
	MemberID    string `json:"memberID,omitempty"`    //成员id
	MemberType  int64  `json:"memberType,omitempty"`  //成员类型:1:设备 2:用户
	ProductID   string `json:"productID,omitempty"`   //产品id 获取产品id下的所有设备信息
	DeviceName  string `json:"deviceName,omitempty"`  //设备名
	ActionId    string `json:"actionId,omitempty"`    //产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	InputParams string `json:"inputParams,omitempty"` //输入参数
}

type SendDcActionResp struct {
	ClientToken  string `json:"clientToken,omitempty"`  //调用id
	OutputParams string `json:"outputParams,omitempty"` //输出参数 注意：此字段可能返回 null，表示取不到有效值。
	Status       string `json:"status,omitempty"`       //返回状态
	Code         int64  `json:"code,string,omitempty"`  //设备返回状态码
}

type SendDcPropertyReq struct {
	MemberID      string `json:"memberID,omitempty"`                      //成员id
	MemberType    int64  `json:"memberType,omitempty"`                    //成员类型:1:设备 2:用户
	ProductID     string `json:"productID,omitempty"`                     //产品id 获取产品id下的所有设备信息
	DeviceName    string `json:"deviceName,omitempty"`                    //设备名
	Data          string `json:"data,omitempty"`                          //属性数据, JSON格式字符串, 注意字段需要在物模型属性里定义
	DataTimestamp int64  `json:"dataTimestamp,string,omitempty,optional"` //上报数据UNIX时间戳, 仅对Method:reported有效
	Method        string `json:"method,omitempty,optional"`               //请求类型 , 不填该参数或者 desired 表示下发属性给设备, reported 表示模拟设备上报属性
}

type SendDcPropertyResp struct {
	ClientToken string `json:"clientToken,omitempty"` //调用id
	Status      string `json:"status"`                //返回状态
	Code        int64  `json:"code,string"`           //设备返回状态码
	Data        string `json:"data,omitempty"`        //返回信息
}
