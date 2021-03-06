package helper

import (
	"fmt"
	"github.com/balrogsxt/genshin-auto-sign/helper/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var conf *Config = nil

func init() {
	conf = LoadConfig()
}

type Config struct {
	RunMode          string   `yaml:"run_mode"`           //运行模式release=生产环境 test=测试环境 debug=调试环境
	RedirectTokenUrl string   `yaml:"redirect_token_url"` //登录成功后的回调token地址 ,变量%token%
	HttpHost         string   `yaml:"http_host"`          //服务绑定地址
	HttpPort         int      `yaml:"http_port"`          //服务启动端口
	JwtKey           string   `yaml:"jwt_key"`            //jwt密钥
	NewUser          bool     `yaml:"new_user"`           //是否允许新用户使用
	NoRegisterText   string   `yaml:"noreg_text"`         //关闭注册时候的文字说明
	CurlApi          []string `yaml:"curl_api"`           //远程请求API
	Task             []string `yaml:"task"`               //任务corn触发时间
	NotifyImage      struct {
		DomainUrl      string `yaml:"domain_url"`      //可用变量 {NAME} = SaveName的值
		BackgroundFile string `yaml:"background_file"` //背景绘制图片
		FontFile       string `yaml:"font_file"`
		SaveName       string `yaml:"save_name"` //可用变量 {DATE} = 日期Y-m-d
		SavePath       string `yaml:"save_path"` //可用变量 {NAME} = SaveName的值
	} `yaml:"notify_image"`
	Smtp struct {
		Enable   bool   //是否启用
		Host     string //邮件服务器地址
		Port     int    //邮件服务器端口
		User     string //用户
		Password string //密码
		From     string //别名来源
	}
	QQOauth struct {
		ClientId     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
		RedirectUri  string `yaml:"redirect_uri"`
	}
	Mysql struct {
		Host     string
		Port     int
		Name     string
		User     string
		Password string
	}
	Redis struct {
		Host     string
		Port     int
		Password string
		Index    int
	}
	QQBot struct {
		Url               string
		QQ                string
		Key               string
		BindNotifyGroup   []string `yaml:"bind_notify_group"`   //绑定用户成功后通知的群组
		SignNotifyGroup   []string `yaml:"sign_notify_group"`   //签到成功后通知的群组
		ExpireNotifyGroup []string `yaml:"expire_notify_group"` //cookie过期后通知的群组
	}
}

//获取配置文件
func GetConfig() *Config {
	if conf == nil {
		conf = LoadConfig()
	}
	return conf
}

//加载配置文件
func LoadConfig() *Config {
	file, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败: %s", err.Error()))
	}
	conf := Config{}
	if err := yaml.Unmarshal(file, &conf); err != nil {
		panic(fmt.Sprintf("解析配置文件失败: %s", err.Error()))
	}
	log.Info("载入配置文件成功")
	return &conf
}
