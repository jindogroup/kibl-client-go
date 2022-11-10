package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var Env envUtils

type envUtils struct {
}

func (envUtils) GetOrDefault(env, _default string) Value {
	if val := os.Getenv(env); val != "" {
		return Value(val)
	}
	return Value(_default)
}

func (envUtils) Helper(prefix ...string) *_env {
	env := _env{}
	return env.setPrx(prefix...)
}

type _env struct {
	prx string
}

func (e *_env) Prefix() string {
	return e.prx
}

func (e *_env) OrDefault(prx ...string) *_env {
	if len(prx) == 0 || e.prx != "" {
		return e
	} 
	e.prx = strings.Join(prx, "_")
	return e
}

func (e *_env) setPrx(prx ...string) *_env {
	if len(prx) == 0 {
		return e
	}
	e.prx = strings.Trim(strings.Join(prx, "_"), "_")
	return e
}

func (e _env) Get(name string, _default ...interface{}) Value {
	var def string
	if len(_default) > 0 {
		def = fmt.Sprint(_default[0])
	}
	return Env.GetOrDefault(e._prx(name), def)
}

func (e _env) MustGet(name string) (val Value) {
	if val = e.Get(name, ""); !val.Valid() {
		panic("could not find env var " + name)
	}
	return
}

func (e _env) GetString(name string, _default ...interface{}) string {
	return e.Get(name, _default...).String()
}

func (e _env) GetDuration(name string, _default ...interface{}) time.Duration {
	d, err := time.ParseDuration(e.GetString(name, _default...))
	if err != nil {
		return 0
	}
	return d
}

func (e _env) _prx(name string) string {
	return fmt.Sprintf("%s_%s", strings.Trim(e.prx, "_"), name)
}


