package config

import (
	"strings"
	"yoyo-mall/pkg/log"

	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 初始化日志包
	c.initLog()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("./config") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")    // 设置配置文件格式为YAML
	viper.AutomaticEnv()           // 读取匹配的环境变量
	viper.SetEnvPrefix("YOYOMALL") // 读取环境变量的前缀为YOYOMALL
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) initLog() {
	log.Init()
}
