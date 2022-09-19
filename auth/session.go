package auth

import (
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type Session struct {
	result *cognito.AuthenticationResultType
}

func (s *Session) AccessToken() string {
	return *s.result.AccessToken
}

func GetSessionFromConfig(conf *Config) (session *Session, err error) {
	session = new(Session)
	if session.result, err = cognitoLogin(conf.Username, conf.Password, conf.ClientID, conf.Region); err != nil {
		return nil, err
	}
	//todo: implement refresh
	return
}

// GetSession sets the credentials from environment variables
func GetSession(prefix ...string) (session *Session, err error) {
	var conf *Config
	if conf, err = ConfFromEnv(prefix...); err != nil {
		return
	}
	return GetSessionFromConfig(conf)
}

func GetSessionWithCredentials(username, password, clientId, region string) (session *Session, err error) {
	return GetSessionFromConfig(&Config{
		Username: username,
		Password: password,
		ClientID: clientId,
		Region:   region,
	})
}

