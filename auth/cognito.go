package auth

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func cognitoLogin(username, password, clientId, region string) (result *cognito.AuthenticationResultType , err error) {
	conf := &aws.Config{Region: aws.String(region)}
	sess := session.Must(session.NewSession(conf))
	client := cognito.New(sess)

	login := &cognito.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME":    aws.String(username),
			"PASSWORD":    aws.String(password),
		},
		ClientId: aws.String(clientId),
	}
	var resp *cognito.InitiateAuthOutput
	if resp, err = client.InitiateAuth(login); err != nil {
		return
	}
	result = resp.AuthenticationResult
	return
}