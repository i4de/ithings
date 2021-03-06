package productDetail

type AUTH = int64

const (
	AUTH_PWD  AUTH = 1 //账密认证
	AUTH_CERT AUTH = 2 //证书认证
)

type NET = int64

const (
	NET_OTHER NET = 1 //其他
	NET_WIFI  NET = 2 //wi-fi
	NET_G234  NET = 3 //2G/3G/4G
	NET_G5    NET = 4 //5G
	NET_BLE   NET = 5 //蓝牙
	NET_LORA  NET = 6 //LoRaWAN
)

type DEV = int64

const (
	DEV_DEVICE  DEV = 1 //设备
	DEV_GATEWAY DEV = 2 //网关
	DEV_SUBSET  DEV = 3 //子设备
)

type DATA = int64

const (
	DATA_DATA_UNKNOWN DATA = 0 //禁止为此参数
	DATA_CUSTOM       DATA = 1 //自定义
	DATA_TEMPLATE     DATA = 2 //数据模板
)

type AUTO_REG = int64

const (
	AUTO_REG_CLOSE AUTO_REG = 1 //关闭
	AUTO_REG_OPEN  AUTO_REG = 2 //打开
	AUTO_REG_AUTO  AUTO_REG = 3 //打开并自动创建设备
)
