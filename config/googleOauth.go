package config

import "errors"

type GoogleOAuth struct {
	ClientId string
}

func NewGoogleOauth() (*GoogleOAuth, error) {

	clientId := getenvDefault("GOOGLE_OAUTH_CLIENT_ID", "")
	if len(clientId) == 0 {
		return nil, errors.New("google oauth config error : clientId is empty")
	}

	return &GoogleOAuth{
		ClientId: clientId,
	}, nil
}
