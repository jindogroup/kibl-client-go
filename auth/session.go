package auth

import (
	"time"

	"github.com/sirupsen/logrus"
)

type secureString string

func (s secureString) String() string {
	return "[redacted]"
}

func (s secureString) value() string {
	return string(s)
}


type Session interface {
	SetLogger(log *logrus.Entry) Session
	Authenticate() error
	AccessToken() string
	RefreshToken() string
	Refresh(token string) error
	ExpiresIn() time.Duration
}

func GetSessionFromConfig(conf *Config) (session Session, err error) {
	session = Cognito(conf).SetLogger(Logger())
	if err = session.Authenticate(); err != nil {
		return nil, err
	}
	return
}

func GetSessionFromConfigWithLogger(logger *logrus.Entry, conf *Config) (session Session, err error) {
	session = Cognito(conf).SetLogger(logger)
	if err = session.Authenticate(); err != nil {
		return nil, err
	}
	return
}

// GetSession sets the credentials from environment variables
func GetSession(logger *logrus.Entry, prefix ...string) (session Session, err error) {
	var conf *Config
	if conf, err = ConfFromEnv(prefix...); err != nil {
		return
	}
	return GetSessionFromConfigWithLogger(logger, conf)
}

func GetSessionWithCredentials(log *logrus.Entry, username, password, clientId, region string) (session Session, err error) {
	return GetSessionFromConfig(&Config{
		Username: username,
		Password: password,
		ClientID: clientId,
		Region:   region,
	})
}

