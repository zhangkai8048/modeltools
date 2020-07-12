package conf

// model保存路径
const ModelPath = "./models/"
// 是否覆盖已存在model
const ModelReplace = true
// 数据库驱动
const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   string
	User   string
	Pwd    string
	DbName string
}
// 数据库链接配置
var MasterDbConfig DbConf = DbConf{
	Host:   "39.107.82.101",
	Port:   "3306",
	User:   "juran",
	Pwd:    "JrGJ@0325#@!",
	DbName: "juranwuyou-user",
}
