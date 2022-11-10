package utils

import (
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type Value string

func (v Value) Int() (i int) {
	i, _ = strconv.Atoi(strings.Trim(string(v), " "))
	return 
}

func (v Value) String() string {
	return string(v)
}

func (v Value) Bool() (b bool) {
	var e error
	if b, e = strconv.ParseBool(v.String()); e != nil {
		log.WithError(e).Errorf("error parsing %s", v.String())
	}
	return 
}

func (v Value) Duration() time.Duration {
	d, e := time.ParseDuration(v.String())
	if e != nil {
		log.Errorf("Error parsing duration '%v'. %v", v, v, e)
	}
	return d
}

func (v Value) Valid() bool {
	return strings.Trim(string(v), " ") != ""
}