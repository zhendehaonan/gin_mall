package conf

//ini包下载：go get gopkg.in/ini.v1
import (
	"gin_mall/dao"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppModel string
	HttpPort string

	DB         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisDb       string
	RedisAddr     string
	RedisPassword string
	RedisDbName   string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtPass    string

	Host        string
	ProductPath string
	AvatarPAth  string
)

// 初始化配置文件
func Init() {
	//本地读取环境变量
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadMysql(file)
	LoadRedis(file)
	LoadEmail(file)
	LoadPath(file)
	//mysql 读 主库
	pathRead := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	//mysql 写 从库
	pathWrite := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	dao.Database(pathRead, pathWrite)
}
func LoadServer(file *ini.File) {
	AppModel = file.Section("service").Key("AppModel").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysql(file *ini.File) {
	DB = file.Section("mysql").Key("DB").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPassword = file.Section("redis").Key("RedisPassword").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtPass = file.Section("email").Key("SmtPass").String()
}

func LoadPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPAth = file.Section("path").Key("AvatarPAth").String()
}
