# kibl-client-go

Golang implementation for the kibl client


# Configuration
Export the following variables for automatic configuration

```bash
export KIBL_USERNAME=appname 
export KIBL_PASSWORD=******** 
export KIBL_CLIENT_ID=clientId 
export KIBL_REGION=us-west-2
```

## Supported Features

The client supports the following

### Get Authenticated Session

You can use this client to obtain an authenticated session for use with kibl public APIs. See the example below.

```golang
  sess, err := auth.GetSession()
	handleError(err)
	req, err := http.NewRequest(http.MethodPost, "https://api.kibl.io/sports/get/info/markets", nil)
	handleError(err)
	req.Header.Set("Authorization", "Bearer " + sess.AccessToken())
	client := &http.Client{}
	resp, err := client.Do(req)
	handleError(err)
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body);
	handleError(err)
	fmt.Println(string(data))
```

You can alternatively get the session using the credentials directly, or alternatively use your own configuration. In these cases, the environment variables are irrelevant
```golang
// using credentials
var session auth.Session = auth.GetSessionWithCredentials(username, password, clientId, region)
// using custom config
session = auth.GetSessionFromConfig(auth.Config{
		Username:      "kibluser",
		Password:      "s0m3p@sswArd",
		ClientID:      "clI3nt1DgoesHere",
		Region:        "us-east-1",
		RefreshBefore: 10*time.Second,
	})
```

It' is also possible to set your own prefixes when setting up the environment variables. 
```golang
conf := auth.ConfFromEnv("MYAPP", "CRED1")
// will create configuration under the hood by reading from the environment like
...
return &Config{
		Username: os.Getenv("MYAPP_CRED1_USERNAME"),
		Password: os.Getenv("MYAPP_CRED1_PASSWORD"),
		ClientID: os.Getenv("MYAPP_CRED1_CLIENT_ID"),
		Region:   os.Getenv("MYAPP_CRED1_REGION"),
	}
...

```

The session will handle token refresh automatically. The token is refreshed _40 seconds_ ahead of the `expiresIn` time set by the token issuer. You can control this behaviour by setting `KIBL_REFRESH_BEFORE` in your environment or setting `auth.Config.RefreshBefore` before the session is created.   

## Todo
- Testing