package auth

type Config struct {
	Username string
	Password string
	ClientID string
	Region   string
}

func ConfFromEnv(prefix ...string) (*Config, error) {
	env := envHelper(prefix...)
	return &Config{
		Username: env.get("USERNAME"),
		Password: env.get("PASSWORD"),
		ClientID: env.get("CLIENT_ID"),
		Region:   env.get("REGION"),
	}, nil
}