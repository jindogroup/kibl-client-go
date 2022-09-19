package auth

import (
	"fmt"
	"os"
	"strings"
)

func envHelper(prefix...string) _env {
	env := _env{}
	env.setPrx(prefix...)
	return env
}

type _env struct {
	prx string
}

func (e *_env) setPrx(prx...string) {
	if len(prx) == 0 {
		e.prx = "KIBL"
		return
	}
	e.prx = strings.Join(prx, "_")
}

func (e _env) get(env string) (value string) {
	var name string = e._prx(env)
	if value = os.Getenv(name); value == "" {
		panic(fmt.Sprintf("%s undefined in environment", name))	
	}
	return
}

func (e _env) _prx(name string) string {
	return fmt.Sprintf("%s_%s", strings.Trim(e.prx, "_"), name)
}