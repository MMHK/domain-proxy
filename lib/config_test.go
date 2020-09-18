package lib

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	confPath := getLocalPath("../tests/config.json")
	conf, err := NewConfig(confPath)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	err = conf.Save()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	t.Log(conf)
}
