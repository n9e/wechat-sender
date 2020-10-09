package config

import (
	"fmt"

	"github.com/toolkits/pkg/file"
)

type Config struct {
	Logger   loggerSection   `yaml:"logger"`
	WeChat   wechatSection   `yaml:"wechat"`
	Consumer consumerSection `yaml:"consumer"`
	Redis    redisSection    `yaml:"redis"`
	HTTP     httpSection     `yaml:"http"`
}

type loggerSection struct {
	Dir       string `yaml:"dir"`
	Level     string `yaml:"level"`
	KeepHours uint   `yaml:"keepHours"`
}

type httpSection struct {
	Listen string `yaml:"listen"`
}

type redisSection struct {
	Addr    string         `yaml:"addr"`
	Pass    string         `yaml:"pass"`
	DB      int            `yaml:"db"`
	Idle    int            `yaml:"idle"`
	Timeout timeoutSection `yaml:"timeout"`
}

type timeoutSection struct {
	Conn  int `yaml:"conn"`
	Read  int `yaml:"read"`
	Write int `yaml:"write"`
}

type consumerSection struct {
	Enable bool   `yaml:"enable"`
	Queue  string `yaml:"queue"`
	Worker int    `yaml:"worker"`
}

type wechatSection struct {
	CorpID  string `yaml:"corp_id"`
	AgentID int    `yaml:"agent_id"`
	Secret  string `yaml:"secret"`
}

var yaml Config

func Get() Config {
	return yaml
}

func ParseConfig(yf string) error {
	err := file.ReadYaml(yf, &yaml)
	if err != nil {
		return fmt.Errorf("cannot read yml[%s]: %v", yf, err)
	}
	return nil
}
