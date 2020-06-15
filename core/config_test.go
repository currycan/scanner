package core

import (
	"fmt"
	"testing"
)

var cfgFile string
var c Config

func TestParseConf(t *testing.T) {
	cfgFile = "../example/scanner.yaml"
	c.Name = cfgFile
	if err := c.LoadConfigFromYaml(); err != nil {
		t.Fatalf("%s: %s", t.Name(), err.Error())
	}
}

func TestWatchConfig(t *testing.T) {
	cfgFile = "../example/scanner.yaml"
	c.Name = cfgFile
	if err := c.LoadConfigFromYaml(); err != nil {
		t.Fatalf("%s: %s", t.Name(), err.Error())
	}
	var configChange = make(chan int, 1)
	if err := c.WatchConfig(configChange); err != nil {
		t.Fatalf("%s: %s", t.Name(), err.Error())
	}
}

func TestUnmarshalStruct(t *testing.T) {
	cfgFile = "../example/scanner.yaml"
	c.Name = cfgFile
	if err, s := c.UnmarshalStruct(); err != nil {
		t.Fatalf("%s: %s", t.Name(), err.Error())
	} else {
		fmt.Printf("service is %v", s)
	}
}
