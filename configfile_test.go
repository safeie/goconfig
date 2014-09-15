package goconfig

import (
	"testing"
)

func TestConfig1(t *testing.T) {
	var f string = "example.conf"
	c, err := ReadConfigFile(f)
	if err != nil {
		t.Error(err.Error())
		t.Error("read config faild!")
	}
	
	// section
	_, err = c.GetSection("log")
	if err != nil {
		t.Error(err.Error())
	}

	// string
	sv, err := c.GetString("redis", "redisAddr")
	if err != nil {
		t.Error(err.Error())
	}
	if sv != "192.168.1.80:6379" {
		t.Error("c.GetString(\"redis\", \"redisAddr\") should be 192.168.1.80:6379, but: ", sv)
	}

	// int
	iv, err := c.GetInt64("log", "logDays")
	if err != nil {
		t.Error(err.Error())
	}
	if iv != 14 {
		t.Error("c.GetString(\"log\", \"logDays\") should be 14, but: ", iv)
	}

	// float
	fv, err := c.GetFloat("log", "logSize")
	if err != nil {
		t.Error(err.Error())
	}
	if fv != 1.5 {
		t.Error("c.GetString(\"log\", \"logSize\") should be 1.5, but: ", fv)
	}

	// bool
	bv, err := c.GetBool("log", "logOpen")
	if err != nil {
		t.Error(err.Error())
	}
	if bv != false {
		t.Error("c.GetString(\"log\", \"logOpen\") should be true, but: ", bv)
	}

}
