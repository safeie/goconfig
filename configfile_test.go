package goconfig

import "testing"
import "os"

func TestConfig1(t *testing.T) {
	var f string = "example.conf"
	c, err := ReadConfigFile(f)
	if err != nil {
		t.Error(err.Error())
		t.Error("read config faild!")
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
	iv, err := c.GetInt("log", "logDays")
	if err != nil {
		t.Error(err.Error())
	}
	if iv != 14 {
		t.Error("c.GetInt(\"log\", \"logDays\") should be 14, but: ", iv)
	}

	// int64
	iv64, err := c.GetInt64("log", "logDays")
	if err != nil {
		t.Error(err.Error())
	}
	if iv64 != 14 {
		t.Error("c.GetInt64(\"log\", \"logDays\") should be 14, but: ", iv64)
	}

	// float
	fv, err := c.GetFloat("log", "logSize")
	if err != nil {
		t.Error(err.Error())
	}
	if fv != 1.5 {
		t.Error("c.GetFloat(\"log\", \"logSize\") should be 1.5, but: ", fv)
	}

	// bool
	bv, err := c.GetBool("log", "logOpen")
	if err != nil {
		t.Error(err.Error())
	}
	if bv != false {
		t.Error("c.GetString(\"log\", \"logOpen\") should be true, but: ", bv)
	}

	// MustString not set
	sv = c.MustString("redis", "redisAddrNo", "127.0.0.0")
	if sv != "127.0.0.0" {
		t.Error("c.MustString(\"redis\", \"redisAddrNo\") should be 127.0.0.0, but: ", sv)
	}

	// MustString yet set
	sv = c.MustString("redis", "redisAddr", "127.0.0.0")
	if sv != "192.168.1.80:6379" {
		t.Error("c.MustString(\"redis\", \"redisAddr\") should be 192.168.1.80:6379, but: ", sv)
	}

	// MustInt not set
	iv = c.MustInt("log", "logDaysNo", 7)
	if iv != 7 {
		t.Error("c.MustInt(\"log\", \"logDaysNo\") should be 7, but: ", iv)
	}

	// MustInt yet set
	iv = c.MustInt("log", "logDays", 7)
	if iv != 14 {
		t.Error("c.MustInt(\"log\", \"logDays\") should be 14, but: ", iv)
	}

	// MustInt64 not set
	iv64 = c.MustInt64("log", "logDaysNo", 7)
	if iv64 != 7 {
		t.Error("c.MustInt64(\"log\", \"logDaysNo\") should be 7, but: ", iv64)
	}

	// MustInt64 yet set
	iv64 = c.MustInt64("log", "logDays", 7)
	if iv64 != 14 {
		t.Error("c.MustInt64(\"log\", \"logDays\") should be 14, but: ", iv64)
	}

	// MustFloat not set
	fv = c.MustFloat("log", "logSizeNo", 3.8)
	if fv != 3.8 {
		t.Error("c.MustFloat(\"log\", \"logSizeNo\") should be 3.8, but: ", fv)
	}

	// MustFloat yet set
	fv = c.MustFloat("log", "logSize", 3.8)
	if fv != 1.5 {
		t.Error("c.MustFloat(\"log\", \"logSize\") should be 1.5, but: ", fv)
	}

	// MustBool not set
	bv = c.MustBool("log", "logOpenNo", true)
	if bv != true {
		t.Error("c.MustBool(\"log\", \"logOpenNo\") should be true, but: ", bv)
	}

	// MustBool yet set
	bv = c.MustBool("log", "logOpen", true)
	if bv != false {
		t.Error("c.MustBool(\"log\", \"logOpen\") should be false, but: ", bv)
	}

	// Env test
	os.Setenv("GOCONFIG_ENV_TEST", "foo")
	os.Setenv("GOCONFIG_ENV_TEST2", "bar")
	c.AddOption("env", "test", parseEnv("xx{{ENV:GOCONFIG_ENV_TEST}}xx{{ ENV:GOCONFIG_ENV_TEST2 }}"))
	ev := c.MustString("env", "test", "")
	if ev != "xxfooxxbar" {
		t.Error("c.MustBool(\"env\", \"test\") should be xxfooxxbar, but: ", ev)
	}

	// variable test
	ev = c.MustString("redis", "redisAddr", "")
	if ev != "192.168.1.80:6379" {
		t.Error("c.MustString(\"redis\", \"redisAddr\") should be 192.168.1.80:6379, but: ", ev)
	}
}
