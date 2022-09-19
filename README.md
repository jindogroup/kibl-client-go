# kibl-client-go

Golang implementation for the kibl client


# Usage
Export the following variables for automatic configuration

```bash
export KIBL_USERNAME=appname 
export KIBL_PASSWORD=******** 
export KIBL_CLIENT_ID=clientId 
export KIBL_REGION=us-west-2
```

You can alternatively get the session usig the credentials directly
```golang
auth.GetSessionWithCredentials(username, password, clientId, region)
```

## Supported Features

The client supports the following

### GetAccessTonek
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

## Todo
- Session renewal
- Add logging