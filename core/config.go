package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//读取yaml配置文件
func (c *Config) LoadConfigFromYaml() (err error) {
	if c.Name != "" {
		// 使用 flag 标志中传递的配置文件
		if err, isFile := IsFileExist(c.Name); err == nil && isFile {
			viper.SetConfigFile(c.Name)
		} else if err != nil {
			logrus.Fatalf("error info: %s\n", err)
		} else if !isFile {
			fmt.Printf("%s doesn't exist!\n", c.Name)
			os.Exit(1)
		}
	} else {
		// 从默认配置文件读取
		//logrus.Info("Read config from default path.\n")
		viper.AddConfigPath("./")
		viper.AddConfigPath("/etc/scanner")
		viper.SetConfigName("scanner")
	}
	viper.SetConfigType("yaml")

	// 从环境变量总读取
	viper.AutomaticEnv()
	viper.SetEnvPrefix("scan")
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	return viper.ReadInConfig()
}

//监听配置文件的修改和变动
func (c *Config) WatchConfig(change chan int) (err error) {
	viper.WatchConfig()

	//监听回调函数
	watch := func(e fsnotify.Event) {
		logrus.Warnf("Config file is changed: %s \n", e.String())
		if err := viper.ReadInConfig(); err != nil {
			return
		}
		change <- 1
	}

	viper.OnConfigChange(watch)
	return err
}

//将配置解析为Struct对象
func (c *Config) UnmarshalStruct() (err error, s Services) {
	c.LoadConfigFromYaml()
	viper.Unmarshal(&s)
	return err, s
}
