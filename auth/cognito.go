package auth

import (
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/jindogroup/kibl-client-go/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	cognitoAuthFlowUserPass     = "USER_PASSWORD_AUTH"
	cognitoAuthFlowRefreshToken = "REFRESH_TOKEN_AUTH"
)

func Cognito(config *Config) *CognitoSession {
	s := &CognitoSession{
		clientID: (*secureString)(aws.String(config.ClientID)),
		username: (*secureString)(aws.String(config.Username)),
		password: (*secureString)(aws.String(config.Password)),
		refreshBefore: config.RefreshBefore,
	}
	conf := &aws.Config{Region: aws.String(config.Region)}
	sess := session.Must(session.NewSession(conf))
	s.client = cognito.New(sess)
	return s
}

type CognitoSession struct {
	log *logrus.Entry

	clientID      *secureString
	password      *secureString
	username      *secureString
	refreshBefore time.Duration

	accessToken  *secureString
	refreshToken *secureString
	expiresIn    int64

	refresh *utils.T
	mx      sync.RWMutex

	client *cognito.CognitoIdentityProvider
}

func (s *CognitoSession) SetLogger(log *logrus.Entry) Session {
	s.log = log.WithFields(logrus.Fields{
		"user": s.username.value(), "authenticator": "cognito",
	})
	return s
}

func (s *CognitoSession) Authenticate() (err error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.log.Info("Authenticating user")
	s.accessToken, s.refreshToken = nil, nil
	var res *cognito.InitiateAuthOutput
	if res, err = s.client.InitiateAuth(s.creds(cognitoAuthFlowUserPass, "")); err != nil {
		err = errors.Wrap(err, "USER_PASSWORD authentication failed")
		s.log.WithError(err).Error("Authentication Failed")
		return
	}
	return s.setValidatedResponse(res)
}

func (s *CognitoSession) Refresh(token string) (err error) {
	s.log.Info("Refreshing Session Token")
	var res *cognito.InitiateAuthOutput
	if res, err = s.client.InitiateAuth(s.creds(cognitoAuthFlowRefreshToken, token)); err != nil {
		err = errors.Wrap(err, "Session token refresh failed")
		s.log.WithError(err).Error("Token refresh failed")
		return errors.Wrap(s.Authenticate(), "Session re-authentication failed")
	}
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.setValidatedResponse(res)
}

func (s *CognitoSession) AccessToken() string {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.accessToken.value()
}

func (s *CognitoSession) RefreshToken() string {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.refreshToken.value()
}

func (s *CognitoSession) ExpiresIn() time.Duration {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return time.Duration(s.expiresIn) * time.Second
}

func (s *CognitoSession) setValidatedResponse(res *cognito.InitiateAuthOutput) (err error) {
	if res == nil {
		return errors.New("authentication response was empty")
	} else if res.AuthenticationResult == nil {
		err = errors.New("auth incomplete. More challenge steps required")
		s.log.WithField("challenge", logrus.Fields{
			"name":       res.ChallengeName,
			"parameters": res.ChallengeParameters,
			"session":    res.Session,
		}).WithError(err).Error()
		return
	}
	return s.setResponse(res)
}

func (s *CognitoSession) setResponse(res *cognito.InitiateAuthOutput) (err error) {
	if s.refreshToken == nil {
		s.log.Info("Session login successfull !!")
		s.refreshToken = (*secureString)(aws.String(*res.AuthenticationResult.RefreshToken))
	} else {
		s.log.Info("Session token refresh successfull !!")
	}

	s.accessToken = (*secureString)(aws.String(*res.AuthenticationResult.AccessToken))
	s.expiresIn = *res.AuthenticationResult.ExpiresIn

	if s.refresh != nil {
		s.refresh.Cancel()
	}

	s.refresh = utils.Timer()
	refreshTime := (time.Duration(s.expiresIn) * time.Second) - s.refreshBefore

	go func(refreshTime time.Duration, refresh <-chan bool) {
		s.log.Infof("Will refresh session token in %v", refreshTime.Round(time.Second))
		if ok := <-refresh; ok {
			s.Refresh(s.refreshToken.value())
		}
	}(refreshTime, s.refresh.After(refreshTime))
	return
}

func (s *CognitoSession) creds(authFlow, token string) (cr *cognito.InitiateAuthInput) {
	cr = &cognito.InitiateAuthInput{
		AuthFlow: aws.String(authFlow), ClientId: aws.String(s.clientID.value()),
	}
	switch authFlow {
	case cognitoAuthFlowUserPass:
		cr.AuthParameters = map[string]*string{
			"USERNAME": aws.String(s.username.value()),
			"PASSWORD": aws.String(s.password.value()),
		}
	case cognitoAuthFlowRefreshToken:
		cr.AuthParameters = map[string]*string{
			"REFRESH_TOKEN": aws.String(token),
		}
	}
	return
}
