package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"encoding/json"

	"gopkg.in/yaml.v2"
)

//yaml文件中map Nginx字段设置
type Nginx struct {
	Port    int    `yaml:"Port"`
	LogPath string `yaml:"LogPath"`
	Path    string `yaml:"Path"`
}

type MQ struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password`
}
type HTTP struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}
type Test struct {
	User []string `yaml:"user"` //这个列别总是出错
	MQTT MQ       `yaml:"mqtt`
	Http HTTP     `yaml:"http"`
}

type TestUrl struct {
	Name      string `yaml:"name"`
	Url       string `yaml:"url"`
	ProxyPort string `yaml:"proxy_port,omitempty"`
}

//Config 系统配置配置 yaml 整个字段配置
type Config struct {
	Name      string    `yaml:"SiteName"`
	Addr      string    `yaml:"SiteAddr"`
	HTTPS     bool      `yaml:"Https"`
	SiteNginx Nginx     `yaml:"Nginx"`
	UserTest  Test      `yaml:"test"`
	TestUrls  []TestUrl `yaml:"test_url"`
}

func main() {
	var setting Config
	config, err := ioutil.ReadFile("./first.yaml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &setting)
	fmt.Println(setting.Name)
	fmt.Println(setting.Addr)
	fmt.Println(setting.HTTPS)

	fmt.Println(setting.SiteNginx.Port)
	fmt.Println(setting.SiteNginx.LogPath)
	fmt.Println(setting.SiteNginx.Path)

	//列表
	fmt.Println(setting.UserTest.User)
	fmt.Println(setting.UserTest.User[0])
	//列表+map
	fmt.Println(setting.TestUrls[0].Name)
	fmt.Println(setting.TestUrls[1].Name)

	conf, err := ReadYamlConfig("./first.yaml")
	if err != nil {
		fmt.Print(err)
	}
	byts, err := json.Marshal(conf) //byts字符数组类型
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(byts))

}

//read yaml config
//注：path为yaml或yml文件的路径
func ReadYamlConfig(path string) (*Config, error) {
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	return conf, nil
}
