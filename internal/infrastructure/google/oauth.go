package google

import (
	"context"
	"errors"

	"cloud.google.com/go/auth/credentials/idtoken"
)

type OAuthProvider struct {
	clientId string
}

type OAuthProviderParams struct {
	ClientId string
}

func (p *OAuthProviderParams) validate() error {
	if len(p.ClientId) == 0 {
		return errors.New("google oauth provider error : invalid credentials")
	}
	return nil
}

func NewOAuthProvider(params OAuthProviderParams) (*OAuthProvider, error) {

	if err := params.validate(); err != nil {
		return nil, err
	}

	return &OAuthProvider{
		clientId: params.ClientId,
	}, nil
}

// func (p *OAuthProvider) ClientId() string    { return p.clientId }
// func (p *OAuthProvider) Secret() string      { return p.secret }
// func (p *OAuthProvider) RedirectUrl() string { return p.redirectUrl }

// ValidateToken checks if the token is valid and returns payload's claims.
func (p *OAuthProvider) ValidateToken(ctx context.Context, idToken string) (map[string]any, error) {

	payload, err := idtoken.Validate(ctx, idToken, p.clientId)
	if err != nil {
		return nil, err
	}
	return payload.Claims, nil
}
