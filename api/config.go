package api

import (
	"time"

	"github.com/jindogroup/kibl-client-go/utils"
)

type Conf struct {
	InboundHost    string
	OutboundHost   string
	Host           string
	RequestTimeout time.Duration
	RetryTimeout   time.Duration
	RetryTimes     int
	DisableAuth    bool

	prefix string
}

func Config(prefix ...string) (conf *Conf) {
	env := utils.Env.Helper(prefix...).OrDefault("KIBL_API")

	conf = &Conf{
		Host:           env.GetString("URL", "https://api.kibl.io"),
		RetryTimeout:   env.Get("RETRY_TIMEOUT", "5s").Duration(),
		RequestTimeout: env.Get("RETRY_TIMEOUT", "15s").Duration(),
		RetryTimes:     env.Get("RETRY_TIMES", "3").Int(),
		DisableAuth:    env.Get("DISABLE_AUTH", "false").Bool(),
		prefix:         env.Prefix(),
	}
	conf.InboundHost = env.GetString("INBOUND_URL", conf.Host)
	conf.OutboundHost = env.GetString("OUTOUND_URL", conf.Host)
	return
}
