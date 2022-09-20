package auth

import "time"

type Config struct {
	Username      string
	Password      string
	ClientID      string
	Region        string
	RefreshBefore time.Duration
}

func ConfFromEnv(prefix ...string) (conf *Config, err error) {
	env := envHelper(prefix...)
	conf = &Config{
		Username:      env.get("USERNAME", false),
		Password:      env.get("PASSWORD", false),
		ClientID:      env.get("CLIENT_ID", false),
		Region:        env.get("REGION", false),
		RefreshBefore: 40 * time.Second,
	}

	if refresh := env.get("REFRESH_BEFORE", true); refresh == "" {
		return 
	} else if duration, err := time.ParseDuration(refresh); err == nil {
		conf.RefreshBefore = duration
	}

	return 
}
